package tinvest_client

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"tinvest-go/internal/model"
	contractv1 "tinvest-go/internal/pb"
	grpc_utils "tinvest-go/internal/pkg/grpc"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/pkg/utils"
)

const (
	batchSize = 1000
)

type IClient interface {
	GetAccounts(ctx context.Context, token string) (model.Accounts, error)
	GetFavorites(ctx context.Context, token string) (model.Favorites, error)
	GetPortfolio(ctx context.Context, token string, accountID string) (PortfolioPositions, error)
	GetInstrumentsByIDs(ctx context.Context, token string, IDs []string) (model.Instruments, error)
	GetInstrumentsByTicker(ctx context.Context, token string, ticker string) (model.Instruments, error)
	GetCurrencies(ctx context.Context, token string) (model.Instruments, error)
	GetShares(ctx context.Context, token string) (model.Instruments, error)
	GetBonds(ctx context.Context, token string) (model.Instruments, error)
	GetEtfs(ctx context.Context, token string) (model.Instruments, error)
	GetFutures(ctx context.Context, token string) (model.Instruments, error)
	GetLastPrices(ctx context.Context, token string, IDs []string) (model.LastPrices, error)
	GetOperations(ctx context.Context, token string, accountID string, from time.Time, to time.Time, instrumentIDs []string) (model.Operations, error)
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

func (c *Client) GetInstrumentsByIDs(ctx context.Context, token string, IDs []string) (model.Instruments, error) {
	instruments := make(model.Instruments, len(IDs))
	errGr, errCtx := errgroup.WithContext(ctx)
	mu := sync.Mutex{}
	for _, id := range IDs {
		errGr.Go(func() error {

			req := &contractv1.InstrumentRequest{
				IdType: contractv1.InstrumentIdType_INSTRUMENT_ID_TYPE_UID,
				Id:     id,
			}
			resp, err := c.InstrumentsAPI.GetInstrumentBy(errCtx, req, grpc_utils.NewAuth(token))
			if err != nil {
				logger.Errorf(ctx, "failed to request instrument by id %s: %s", id, err.Error())
				return nil
				// return fmt.Errorf("failed to request instrument by id %s: %w", id, err)
			}

			var currResp *contractv1.CurrencyResponse
			var bondResp *contractv1.BondResponse
			var futureResp *contractv1.FutureResponse
			if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_CURRENCY {
				currResp, err = c.InstrumentsAPI.CurrencyBy(errCtx, req, grpc_utils.NewAuth(token))
				if err != nil {
					return fmt.Errorf("failed to request currency by id %s: %w", id, err)
				}
			} else if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_BOND {
				bondResp, err = c.InstrumentsAPI.BondBy(errCtx, req, grpc_utils.NewAuth(token))
				if err != nil {
					return fmt.Errorf("failed to request bond by id %s: %w", id, err)
				}
			} else if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_FUTURES {
				futureResp, err = c.InstrumentsAPI.FutureBy(errCtx, req, grpc_utils.NewAuth(token))
				if err != nil {
					return fmt.Errorf("failed to request future by id %s: %w", id, err)
				}
			}

			mu.Lock()
			defer mu.Unlock()

			if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_CURRENCY {
				instruments[id] = convertCurrency(currResp.GetInstrument())
			} else if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_BOND {
				instruments[id] = convertBond(bondResp.GetInstrument())
			} else if resp.GetInstrument().GetInstrumentKind() == contractv1.InstrumentType_INSTRUMENT_TYPE_FUTURES {
				instruments[id] = convertFuture(futureResp.GetInstrument())
			} else {
				instruments[id] = convertInstrument(resp.GetInstrument())
			}

			return nil
		})
	}

	err := errGr.Wait()
	if err != nil {
		return nil, err
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
	req := &contractv1.GetLastPricesRequest{
		InstrumentId:  IDs,
		LastPriceType: *contractv1.LastPriceType_LAST_PRICE_EXCHANGE.Enum(),
	}

	resp, err := c.MarketDataAPI.GetLastPrices(ctx, req, grpc_utils.NewAuth(token))
	if err != nil {
		return nil, fmt.Errorf("failed to request last prices: %w", err)
	}

	return convertLastPrices(resp.GetLastPrices()), nil
}

func (c *Client) GetOperations(
	ctx context.Context,
	token string,
	accountID string,
	from time.Time,
	to time.Time,
	instrumentIDs []string,
) (model.Operations, error) {
	errGr, errCtx := errgroup.WithContext(ctx)
	mu := sync.Mutex{}

	res := model.Operations{}
	for _, instrumentID := range instrumentIDs {
		errGr.Go(func() error {
			req := &contractv1.GetOperationsByCursorRequest{
				AccountId:          accountID,
				InstrumentId:       utils.Ptr(instrumentID),
				Limit:              utils.Ptr(int32(batchSize)),
				State:              contractv1.OperationState_OPERATION_STATE_EXECUTED.Enum(),
				WithoutCommissions: utils.Ptr(true),
				WithoutOvernights:  utils.Ptr(false),
				WithoutTrades:      utils.Ptr(true),
				From:               timestamppb.New(from),
				To:                 timestamppb.New(to),
			}

			operations := model.Operations{}
			for {
				resp, err := c.OperationsAPI.GetOperationsByCursor(errCtx, req, grpc_utils.NewAuth(token))
				if err != nil {
					logger.Errorf(ctx, "failed to request operations for instrument %s: %s", instrumentID, err.Error())
					return nil
					// return fmt.Errorf("failed to request operations for instrument %s: %w", instrumentID, err)
				}

				operations = append(operations, convertOperations(resp.GetItems())...)

				if !resp.GetHasNext() {
					break
				}
				req.Cursor = utils.Ptr(resp.GetNextCursor())
			}

			mu.Lock()
			defer mu.Unlock()
			res = append(res, operations...)

			return nil
		})
	}

	err := errGr.Wait()
	if err != nil {
		return nil, err
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
