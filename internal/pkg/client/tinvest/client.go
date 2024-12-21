package tinvest_client

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"tinvest-go/internal/model"
	contractv1 "tinvest-go/internal/pb"
	"tinvest-go/internal/pkg/cache"
	grpc_utils "tinvest-go/internal/pkg/grpc"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/pkg/utils"
)

const (
	batchSize1k = 1000
	batchSize3k = 3000

	headerXRateLimitRemaining = "x-ratelimit-remaining"
	headerXRateLimitReset     = "x-ratelimit-reset"
)

type IClient interface {
	GetAccounts(ctx context.Context, token string) (model.Accounts, error)
	GetFavorites(ctx context.Context, token string) (model.Favorites, error)
	GetInstruments(ctx context.Context, token string) (model.Instruments, error)
	GetInstrumentsByIDs(ctx context.Context, token string, IDs []string) (model.Instruments, error)
	GetInstrumentsByTicker(ctx context.Context, token string, ticker string) (model.Instruments, error)
	GetCurrencies(ctx context.Context, token string) (model.Instruments, error)
	GetShares(ctx context.Context, token string) (model.Instruments, error)
	GetBonds(ctx context.Context, token string) (model.Instruments, error)
	GetEtfs(ctx context.Context, token string) (model.Instruments, error)
	GetFutures(ctx context.Context, token string) (model.Instruments, error)
	GetLastPrices(ctx context.Context, token string, IDs []string) (model.LastPrices, error)
	GetPortfolio(ctx context.Context, token string, accountID string) (PortfolioPositions, error)
	GetOperations(ctx context.Context, token string, accountID string, from time.Time, to time.Time) (model.Operations, error)
	GetOperationsByInstrumentID(ctx context.Context, token string, accountID string, from time.Time, to time.Time, instrumentIDs []string) (model.Operations, error)
	GetCandles(ctx context.Context, token string, instrumentID string, interval contractv1.CandleInterval, from time.Time, to time.Time) (model.Candles, error)
}

type Client struct {
	InstrumentsAPI contractv1.InstrumentsServiceClient
	MarketDataAPI  contractv1.MarketDataServiceClient
	OperationsAPI  contractv1.OperationsServiceClient
	OrdersAPI      contractv1.OrdersServiceClient
	SignalsAPI     contractv1.SignalServiceClient
	UsersAPI       contractv1.UsersServiceClient
}

const (
	ProdHost    = "invest-public-api.tinkoff.ru:443"
	SandBoxHost = "sandbox-invest-public-api.tinkoff.ru:443"
)

func NewClient(conn grpc.ClientConnInterface) IClient {
	return &Client{
		InstrumentsAPI: contractv1.NewInstrumentsServiceClient(conn),
		MarketDataAPI:  contractv1.NewMarketDataServiceClient(conn),
		OperationsAPI:  contractv1.NewOperationsServiceClient(conn),
		OrdersAPI:      contractv1.NewOrdersServiceClient(conn),
		SignalsAPI:     contractv1.NewSignalServiceClient(conn),
		UsersAPI:       contractv1.NewUsersServiceClient(conn),
	}
}

func (c *Client) GetAccounts(ctx context.Context, token string) (model.Accounts, error) {
	resp, err := c.UsersAPI.GetAccounts(ctx, &contractv1.GetAccountsRequest{}, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request accounts: %w", err)
	}

	return convertAccounts(resp.GetAccounts()), nil
}

func (c *Client) GetFavorites(ctx context.Context, token string) (model.Favorites, error) {
	req := &contractv1.GetFavoritesRequest{}
	resp, err := c.InstrumentsAPI.GetFavorites(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request favorites: %w", err)
	}

	return convertFavorites(resp.GetFavoriteInstruments()), nil
}

func (c *Client) GetInstrumentsByIDs(ctx context.Context, token string, IDs []string) (model.Instruments, error) {
	instruments, err := c.GetInstruments(ctx, token)
	if err != nil {
		return nil, err
	}

	ids := utils.ToSet(IDs, func(id string) string { return id })

	res := make(model.Instruments, 0, len(IDs))
	for _, instrument := range instruments {
		if ids.Has(instrument.ID) {
			res = append(res, instrument)
		}
	}

	return res, nil
}

func (c *Client) GetInstrumentsByIDsSlow(ctx context.Context, token string, IDs []string) (model.Instruments, error) {
	instruments := make(model.Instruments, 0, len(IDs))
	var header metadata.MD
	for _, id := range IDs {
		if id == "" {
			continue
		}

		wait(header)
		req := &contractv1.InstrumentRequest{
			IdType: contractv1.InstrumentIdType_INSTRUMENT_ID_TYPE_UID,
			Id:     id,
		}
		resp, err := c.InstrumentsAPI.GetInstrumentBy(ctx, req, grpc_utils.NewAuth(token), grpc.Header(&header))
		if err != nil {
			if status, ok := status.FromError(err); ok {
				switch status.Code() {
				case codes.NotFound:
					logger.Errorf(ctx, "failed to request instrument by id %s - not found: %s", id, err.Error())
					continue
				case codes.InvalidArgument:
					logger.Errorf(ctx, "failed to request instrument by id %s - invalid argument: %s", id, err.Error())
					continue
				}
			}
			return nil, fmt.Errorf("failed to request instrument by id %s: %w", id, err)
		}

		var currResp *contractv1.CurrencyResponse
		var bondResp *contractv1.BondResponse
		var futureResp *contractv1.FutureResponse
		if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_CURRENCY {
			currResp, err = c.InstrumentsAPI.CurrencyBy(ctx, req, grpc_utils.NewAuth(token))
			if err != nil {
				return nil, fmt.Errorf("failed to request currency by id %s: %w", id, err)
			}
			instruments = append(instruments, convertCurrency(currResp.GetInstrument()))
		} else if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_BOND {
			bondResp, err = c.InstrumentsAPI.BondBy(ctx, req, grpc_utils.NewAuth(token))
			if err != nil {
				return nil, fmt.Errorf("failed to request bond by id %s: %w", id, err)
			}
			instruments = append(instruments, convertBond(bondResp.GetInstrument()))
		} else if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_FUTURES {
			futureResp, err = c.InstrumentsAPI.FutureBy(ctx, req, grpc_utils.NewAuth(token))
			if err != nil {
				return nil, fmt.Errorf("failed to request future by id %s: %w", id, err)
			}
			instruments = append(instruments, convertFuture(futureResp.GetInstrument()))
		} else {
			instruments = append(instruments, convertInstrument(resp.GetInstrument()))
		}
	}

	return instruments, nil
}

func (c *Client) GetInstrumentsByTicker(ctx context.Context, token string, ticker string) (model.Instruments, error) {
	req := &contractv1.FindInstrumentRequest{
		Query: strings.ToUpper(ticker),
	}
	resp, err := c.InstrumentsAPI.FindInstrument(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request to fine instrument by ticker %s: %w", ticker, err)
	}
	if len(resp.GetInstruments()) == 0 {
		return nil, nil
	}

	instrumentsIDs := make([]string, 0, len(resp.GetInstruments()))
	for _, instrument := range resp.GetInstruments() {
		instrumentsIDs = append(instrumentsIDs, instrument.GetUid())
	}

	return c.GetInstrumentsByIDs(ctx, token, instrumentsIDs)
}

func (c *Client) GetInstruments(ctx context.Context, token string) (model.Instruments, error) {
	instrumentsIf, exists := cache.Get("tinvestClient.GetInstruments")
	if exists {
		return instrumentsIf.(model.Instruments), nil
	}

	var currencies model.Instruments
	var shares model.Instruments
	var bonds model.Instruments
	var etfs model.Instruments
	var futures model.Instruments
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var err error
		currencies, err = c.GetCurrencies(egCtx, token)
		if err != nil {
			return fmt.Errorf("failed to get currencies: %w", err)
		}
		return nil
	})
	eg.Go(func() error {
		var err error
		shares, err = c.GetShares(ctx, token)
		if err != nil {
			return fmt.Errorf("failed to get shares: %w", err)
		}
		return nil
	})
	eg.Go(func() error {
		var err error
		bonds, err = c.GetBonds(ctx, token)
		if err != nil {
			return fmt.Errorf("failed to get bonds: %w", err)
		}
		return nil
	})
	eg.Go(func() error {
		var err error
		etfs, err = c.GetEtfs(ctx, token)
		if err != nil {
			return fmt.Errorf("failed to get etfs: %w", err)
		}
		return nil
	})
	eg.Go(func() error {
		var err error
		futures, err = c.GetFutures(ctx, token)
		if err != nil {
			return fmt.Errorf("failed to get futures: %w", err)
		}
		return nil
	})

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	instruments := make(model.Instruments, 0, len(currencies)+len(shares)+len(bonds)+len(etfs)+len(futures))
	instruments = append(instruments, currencies...)
	instruments = append(instruments, shares...)
	instruments = append(instruments, bonds...)
	instruments = append(instruments, etfs...)
	instruments = append(instruments, futures...)

	// Когда исин равен тикеру или фиги равен тикеру или тикер заканчивается на -RM, то ищем оригинальный тикер
	for _, instrument := range instruments {
		if instrument.Isin == instrument.Ticker ||
			instrument.Figi == instrument.Ticker ||
			strings.HasSuffix(instrument.Ticker, "-RM") {
			for _, origInstrument := range instruments {
				if origInstrument.Isin == instrument.Isin &&
					origInstrument.Ticker != instrument.Ticker &&
					origInstrument.Figi != origInstrument.Ticker &&
					!strings.HasSuffix(origInstrument.Ticker, "-RM") &&
					strings.HasPrefix(origInstrument.Figi, "BBG") {
					instrument.OriginalID = origInstrument.ID
					break
				}
			}
			if instrument.OriginalID != "" {
				continue
			}
			for _, origInstrument := range instruments {
				if origInstrument.Isin == instrument.Isin &&
					origInstrument.Ticker != instrument.Ticker &&
					origInstrument.Figi != origInstrument.Ticker &&
					!strings.HasSuffix(origInstrument.Ticker, "-RM") {
					instrument.OriginalID = origInstrument.ID
					break
				}
			}
		}
	}

	// Когда фиги начинается с TCSXX, то попробуем найти инструмент с таким же ISIN,
	// но с фиги начинающейся TCS00
	for _, instrument := range instruments {
		if instrument.OriginalID != "" {
			continue
		}
		if strings.HasPrefix(instrument.Figi, "TCS") &&
			!strings.HasPrefix(instrument.Figi, "TCS00") {
			for _, origInstrument := range instruments {
				if origInstrument.Isin == instrument.Isin && strings.HasPrefix(origInstrument.Figi, "TCS00") {
					instrument.OriginalID = origInstrument.ID
					break
				}
			}
		}
	}

	// Когда тикер заканчивается на -RM, то ищем оригинальный тикер
	for _, instrument := range instruments {
		if instrument.OriginalID != "" {
			continue
		}
		if strings.HasSuffix(instrument.Ticker, "-RM") {
			for _, origInstrument := range instruments {
				if instrument.Isin == origInstrument.Isin &&
					instrument.Ticker != origInstrument.Ticker &&
					origInstrument.Ticker != origInstrument.Isin {
					instrument.OriginalID = origInstrument.ID
					break
				}
			}
		}
	}

	// Когда исин равен тикеру, то ищем оригинальный тикер
	for _, instrument := range instruments {
		if instrument.OriginalID != "" {
			continue
		}
		if instrument.Isin == instrument.Ticker {
			for _, origInstrument := range instruments {
				if instrument.Isin == origInstrument.Isin &&
					origInstrument.Isin != origInstrument.Ticker &&
					!strings.HasSuffix(origInstrument.Ticker, "-RM") {
					instrument.OriginalID = origInstrument.ID
					break
				}
			}
		}
	}

	return instruments, nil
}

func (c *Client) GetCurrencies(ctx context.Context, token string) (model.Instruments, error) {
	req := &contractv1.InstrumentsRequest{
		InstrumentStatus: contractv1.InstrumentStatus_INSTRUMENT_STATUS_ALL.Enum(),
	}

	resp, err := c.InstrumentsAPI.Currencies(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request currencies: %w", err)
	}

	return convertCurrencies(resp.GetInstruments()), nil
}

func (c *Client) GetShares(ctx context.Context, token string) (model.Instruments, error) {
	req := &contractv1.InstrumentsRequest{
		InstrumentStatus: contractv1.InstrumentStatus_INSTRUMENT_STATUS_ALL.Enum(),
	}

	resp, err := c.InstrumentsAPI.Shares(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request shares: %w", err)
	}

	return convertShares(resp.GetInstruments()), nil
}

func (c *Client) GetBonds(ctx context.Context, token string) (model.Instruments, error) {
	req := &contractv1.InstrumentsRequest{
		InstrumentStatus: contractv1.InstrumentStatus_INSTRUMENT_STATUS_ALL.Enum(),
	}

	resp, err := c.InstrumentsAPI.Bonds(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request bonds: %w", err)
	}

	return convertBonds(resp.GetInstruments()), nil
}

func (c *Client) GetEtfs(ctx context.Context, token string) (model.Instruments, error) {
	req := &contractv1.InstrumentsRequest{
		InstrumentStatus: contractv1.InstrumentStatus_INSTRUMENT_STATUS_ALL.Enum(),
	}

	resp, err := c.InstrumentsAPI.Etfs(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request ETFs: %w", err)
	}

	return convertEtfs(resp.GetInstruments()), nil
}

func (c *Client) GetFutures(ctx context.Context, token string) (model.Instruments, error) {
	req := &contractv1.InstrumentsRequest{
		InstrumentStatus: contractv1.InstrumentStatus_INSTRUMENT_STATUS_ALL.Enum(),
	}

	resp, err := c.InstrumentsAPI.Futures(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request ETFs: %w", err)
	}

	return convertFutures(resp.GetInstruments()), nil
}

func (c *Client) GetLastPrices(ctx context.Context, token string, IDs []string) (model.LastPrices, error) {
	res := make(model.LastPrices, len(IDs))
	chunks := slices.Chunk(IDs, batchSize3k)
	for chunk := range chunks {
		req := &contractv1.GetLastPricesRequest{
			InstrumentId:  chunk,
			LastPriceType: *contractv1.LastPriceType_LAST_PRICE_EXCHANGE.Enum(),
		}

		var header metadata.MD
		resp, err := c.MarketDataAPI.GetLastPrices(ctx, req, grpc_utils.NewAuth(token), grpc.Header(&header))
		if err != nil {
			return nil, fmt.Errorf("failed to request last prices: %w", err)
		}

		for _, lastPrice := range resp.GetLastPrices() {
			res[lastPrice.GetInstrumentUid()] = convertLastPrice(lastPrice)
		}

		wait(header)
	}

	return res, nil
}

func (c *Client) GetPortfolio(ctx context.Context, token string, accountID string) (PortfolioPositions, error) {
	req := &contractv1.PortfolioRequest{
		AccountId: accountID,
	}

	resp, err := c.OperationsAPI.GetPortfolio(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request portfolio: %w", err)
	}

	return convertPortfolioPositions(resp.GetPositions()), nil
}

func (c *Client) GetOperations(
	ctx context.Context,
	token string,
	accountID string,
	from time.Time,
	to time.Time,
) (model.Operations, error) {
	req := &contractv1.GetOperationsByCursorRequest{
		AccountId:          accountID,
		Limit:              utils.Ptr(int32(batchSize1k)),
		State:              contractv1.OperationState_OPERATION_STATE_EXECUTED.Enum(),
		WithoutCommissions: utils.Ptr(true),
		WithoutTrades:      utils.Ptr(true),
		WithoutOvernights:  utils.Ptr(false),
		From:               timestamppb.New(from),
		To:                 timestamppb.New(to),
	}

	res := make(model.Operations, 0, 1000)
	for {
		var header metadata.MD
		resp, err := c.OperationsAPI.GetOperationsByCursor(ctx, req, grpc_utils.NewAuth(token), grpc.Header(&header))
		if err != nil {
			logger.Errorf(ctx, "failed to request operations: %s", err.Error())
			break
			// return nil, fmt.Errorf("failed to request operations for instrument %s: %w", instrumentID, err)
		}

		for _, item := range resp.GetItems() {
			res = append(res, convertOperation(item))
		}

		if !resp.GetHasNext() {
			break
		}

		req.Cursor = utils.Ptr(resp.GetNextCursor())

		wait(header)
	}

	return res, nil
}

func (c *Client) GetOperationsByInstrumentID(
	ctx context.Context,
	token string,
	accountID string,
	from time.Time,
	to time.Time,
	instrumentIDs []string,
) (model.Operations, error) {
	res := make(model.Operations, 0, 100)
	for _, instrumentID := range instrumentIDs {
		req := &contractv1.GetOperationsByCursorRequest{
			AccountId:          accountID,
			InstrumentId:       utils.Ptr(instrumentID),
			Limit:              utils.Ptr(int32(batchSize1k)),
			State:              contractv1.OperationState_OPERATION_STATE_EXECUTED.Enum(),
			WithoutCommissions: utils.Ptr(true),
			WithoutTrades:      utils.Ptr(true),
			WithoutOvernights:  utils.Ptr(false),
			From:               timestamppb.New(from),
			To:                 timestamppb.New(to),
		}

		for {
			var header metadata.MD
			resp, err := c.OperationsAPI.GetOperationsByCursor(ctx, req, grpc_utils.NewAuth(token), grpc.Header(&header))
			if err != nil {
				logger.Errorf(ctx, "failed to request operations for instrument %s: %s", instrumentID, err.Error())
				break
				// return fmt.Errorf("failed to request operations for instrument %s: %w", instrumentID, err)
			}

			for _, item := range resp.GetItems() {
				res = append(res, convertOperation(item))
			}

			if !resp.GetHasNext() {
				break
			}
			req.Cursor = utils.Ptr(resp.GetNextCursor())

			wait(header)
		}
	}

	return res, nil
}

func (c *Client) GetCandles(ctx context.Context, token string, instrumentID string, interval contractv1.CandleInterval, from time.Time, to time.Time) (model.Candles, error) {
	req := &contractv1.GetCandlesRequest{
		InstrumentId: utils.Ptr(instrumentID),
		Interval:     interval,
		From:         timestamppb.New(from),
		To:           timestamppb.New(to),
	}

	resp, err := c.MarketDataAPI.GetCandles(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request candles: %w", err)
	}

	candles := make(model.Candles, 0, len(resp.GetCandles()))
	for _, candle := range resp.GetCandles() {
		candles = append(candles, &model.Candle{
			Time:   candle.GetTime().AsTime(),
			Open:   utils.QuotationToFloat64(candle.GetOpen()),
			Low:    utils.QuotationToFloat64(candle.GetLow()),
			High:   utils.QuotationToFloat64(candle.GetHigh()),
			Close:  utils.QuotationToFloat64(candle.GetClose()),
			Volume: candle.GetVolume(),
		})
	}

	return candles, nil
}

func wait(header metadata.MD) {
	logger.Infof(context.Background(), "try to wait: %+v", header)
	remaining := getRateLimitRemaining(header)
	if remaining == 0 {
		reset := getRateLimitReset(header)
		time.Sleep(time.Duration(reset+1) * time.Second)
	}
}

func getRateLimitRemaining(header metadata.MD) int {
	if rateLimit, ok := header[headerXRateLimitRemaining]; ok {
		if limit, err := strconv.Atoi(rateLimit[0]); err == nil {
			return limit
		}
	}
	return 1
}

func getRateLimitReset(header metadata.MD) int {
	if rateLimitReset, ok := header[headerXRateLimitReset]; ok {
		if reset, err := strconv.Atoi(rateLimitReset[0]); err == nil {
			return reset
		}
	}
	return 0
}
