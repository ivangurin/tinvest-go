package tinvest_service

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"tinvest-go/internal/model"
	contractv1 "tinvest-go/internal/pb"
	tinvest_client "tinvest-go/internal/pkg/client/tinvest"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/pkg/trades"
	exchange_service "tinvest-go/internal/service/exchange"
)

type IService interface {
	GetAccounts(ctx context.Context, token string) (model.Accounts, error)
	GetAccountByID(ctx context.Context, token string, accountID string) (*model.Account, error)
	GetAccountByName(ctx context.Context, token string, name string) (*model.Account, error)
	GetFavorites(ctx context.Context, token string) (model.Instruments, error)
	GetInstruments(ctx context.Context, token string, accountID string, IDs []string) (model.Instruments, error)
	GetInstrumentsByTicker(ctx context.Context, token string, ticker string) (model.Instruments, error)
	GetPortfolio(ctx context.Context, token string, accountID string) (model.Portfolio, error)
	GetPositions(ctx context.Context, token string, accountID string, from time.Time, to time.Time, instrumentIDs []string) (model.Positions, error)
	GetOperations(ctx context.Context, token string, accountID string, from time.Time, to time.Time, instrumentIDs []string) (model.Instruments, model.Operations, error)
	GetTotals(ctx context.Context, token string, accountID string, from time.Time, to time.Time) (model.Totals, error)
	GetCandles(ctx context.Context, token string, instrumentID string, interval contractv1.CandleInterval, from time.Time, to time.Time) (model.Candles, error)
	GetTrades(ctx context.Context, token string, accountID string, from time.Time, to time.Time) (model.Trades, error)
}

type service struct {
	tinvestClient   tinvest_client.IClient
	exchangeService exchange_service.IService
	DrConvTime      time.Time
	DrList          DrList
}

var (
	From = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
)

func NewService(
	tinvestClient tinvest_client.IClient,
	exchangeService exchange_service.IService,
) IService {
	return &service{
		tinvestClient:   tinvestClient,
		exchangeService: exchangeService,
		DrConvTime:      time.Date(2022, 8, 31, 0, 0, 0, 0, time.Local),
		DrList: DrList{
			&DrItem{
				InstrumentTicker:       model.TickerMBT,
				SourceInstrumentTicker: model.TickerMTSS,
				Koeff:                  2,
			},
			&DrItem{
				InstrumentTicker:       model.TickerPHORGS,
				SourceInstrumentTicker: model.TickerPHOR,
				Koeff:                  1.0 / 3.0,
			},
		},
	}
}

func (s *service) GetAccounts(ctx context.Context, token string) (model.Accounts, error) {
	accounts, err := s.tinvestClient.GetAccounts(ctx, token)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *service) GetAccountByID(ctx context.Context, token string, accountID string) (*model.Account, error) {
	accounts, err := s.GetAccounts(ctx, token)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.ID == accountID {
			return account, nil
		}
	}

	return nil, nil
}

func (s *service) GetAccountByName(ctx context.Context, token string, name string) (*model.Account, error) {
	accounts, err := s.GetAccounts(ctx, token)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.Name == name {
			return account, nil
		}
	}

	return nil, fmt.Errorf("account %s not found", name)
}

func (s *service) GetFavorites(ctx context.Context, token string) (model.Instruments, error) {
	favorites, err := s.tinvestClient.GetFavorites(ctx, token)
	if err != nil {
		return nil, err
	}

	res := make(model.Instruments, len(favorites))
	for _, favorite := range favorites {
		instruments, err := s.GetInstruments(ctx, token, "", []string{favorite.ID})
		if err != nil {
			logger.Errorf(ctx, "error on get instrument %s by id %s: %s", favorite.Ticker, favorite.ID, err.Error())
			continue
		}
		if len(instruments) == 0 {
			continue
		}
		res[favorite.ID] = instruments[favorite.ID]
	}

	return res, nil
}

func (s *service) GetInstrumentsByTicker(ctx context.Context, token string, ticker string) (model.Instruments, error) {
	return s.tinvestClient.GetInstrumentsByTicker(ctx, token, ticker)
}

func (s *service) GetInstruments(ctx context.Context, token string, accountID string, IDs []string) (model.Instruments, error) {
	eg, egCtx := errgroup.WithContext(ctx)
	var instruments model.Instruments
	eg.Go(func() error {
		var err error
		instruments, err = s.tinvestClient.GetInstruments(egCtx, token)
		if err != nil {
			return fmt.Errorf("failed to get instruments: %w", err)
		}
		return nil
	})

	var lastPrices model.LastPrices
	eg.Go(func() error {
		var err error
		lastPrices, err = s.GetLastPrices(egCtx, token, accountID, IDs)
		return err
	})

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	// Calculate LastPrice
	for _, instrument := range instruments {
		lastPrice, exists := lastPrices[instrument.ID]
		if exists {
			if instrument.Type == model.InstrumentTypeCurrency {
				if instrument.Ticker == model.TickerRUB {
					instrument.LastPrice = 1
				} else {
					if lastPrice.AbsoluteValue {
						instrument.LastPrice = lastPrice.Value
					} else {
						instrument.LastPrice = lastPrice.Value * float64(instrument.Lot) / instrument.Nominal
					}
				}
			} else if instrument.Type == model.InstrumentTypeBond {
				if lastPrice.AbsoluteValue {
					instrument.LastPrice = lastPrice.Value
				} else {
					instrument.LastPrice = lastPrice.Value / 100 * instrument.Nominal
				}

				// Для фиги TCS00A105146 НКД указан в рублях, хотя везде стоит валюта доллары
				if instrument.Figi == "TCS00A105146" {
					instrument.NKDRub = instrument.NKD
					instrument.NKD, err = s.exchangeService.Convert(ctx, model.CurrencyRUB, instrument.Currency, instrument.NKDRub, time.Now().UTC())
					if err != nil {
						return nil, err
					}
				} else {
					instrument.NKDRub, err = s.exchangeService.Convert(ctx, instrument.Currency, model.CurrencyRUB, instrument.NKD, time.Now().UTC())
					if err != nil {
						return nil, err
					}
				}
			} else if instrument.Type == model.InstrumentTypeFuture {
				if lastPrice.AbsoluteValue {
					instrument.LastPrice = lastPrice.Value
				} else {
					instrument.LastPrice = instrument.LastPrice /
						instrument.MinPriceIncrement *
						instrument.MinPriceIncrementAmount *
						float64(instrument.Lot)
				}
			} else {
				instrument.LastPrice = lastPrice.Value
			}
		}

		instrument.LastPriceRub, err = s.exchangeService.Convert(ctx, instrument.Currency, model.CurrencyRUB, instrument.LastPrice, time.Now().UTC())
		if err != nil {
			return nil, fmt.Errorf("failed to convert currency for share %s: %w", instrument.ID, err)
		}
	}

	return instruments, nil
}

func (s *service) GetPortfolio(ctx context.Context, token string, accountID string) (model.Portfolio, error) {
	positions, err := s.GetPositions(ctx, token, accountID, From, time.Now().UTC(), nil)
	if err != nil {
		return nil, err
	}

	res := make(model.Portfolio, 0, len(positions))
	for _, position := range positions {
		if position.QuantityEnd == 0 {
			continue
		}

		portfolioPosition :=
			&model.PortfolioPosition{
				InstrumentID: position.InstrumentID,
				Name:         position.Name,
				Ticker:       position.Ticker,
				Currency:     position.Currency,
				Quantity:     position.QuantityEnd,
				PriceEnd:     position.PriceEnd,
				PriceEndRub:  position.PriceEndRub,
				ValueEnd:     position.ValueEnd,
				ValueEndRub:  position.ValueEndRub,
			}

		trades := position.Trades.GetUnclosed()
		for _, trade := range trades {
			if trade.QuantitySell == 0 {
				portfolioPosition.Value += trade.ValueBuy
				portfolioPosition.ValueRub += trade.ValueBuyRub
			} else {
				portfolioPosition.Value -= trade.ValueSell
				portfolioPosition.ValueRub -= trade.ValueSellRub
			}
		}

		portfolioPosition.Price = portfolioPosition.Value / portfolioPosition.Quantity
		portfolioPosition.PriceRub = portfolioPosition.ValueRub / portfolioPosition.Quantity

		portfolioPosition.Total = portfolioPosition.ValueEnd - portfolioPosition.Value
		portfolioPosition.TotalRub = portfolioPosition.ValueEndRub - portfolioPosition.ValueRub

		if portfolioPosition.Total != 0 {
			portfolioPosition.Percent = portfolioPosition.Total / portfolioPosition.Value * 100
		}

		if portfolioPosition.TotalRub != 0 {
			portfolioPosition.PercentRub = portfolioPosition.TotalRub / portfolioPosition.ValueRub * 100
		}

		res = append(res, portfolioPosition)
	}

	slices.SortFunc(res, func(a, b *model.PortfolioPosition) int {
		return strings.Compare(a.Ticker, b.Ticker)
	})

	return res, nil
}

func (s *service) GetLastPrices(ctx context.Context, token string, accountID string, instrumentIDs []string) (model.LastPrices, error) {
	lastPrices, err := s.tinvestClient.GetLastPrices(ctx, token, instrumentIDs)
	if err != nil {
		return nil, fmt.Errorf("failed on get last prices: %w", err)
	}

	if accountID != "" {
		portfolioPositions, err := s.tinvestClient.GetPortfolio(ctx, token, accountID)
		if err != nil {
			return nil, fmt.Errorf("failed on get portfolio: %w", err)
		}

		for _, portfolioPosition := range portfolioPositions {
			if portfolioPosition.CurrentPrice == 0 {
				continue
			}

			lastPrice, exists := lastPrices[portfolioPosition.ID]
			if !exists {
				lastPrice = &model.LastPrice{}
				lastPrices[portfolioPosition.ID] = lastPrice
			}

			lastPrice.Value = portfolioPosition.CurrentPrice
			lastPrice.AbsoluteValue = true
		}
	}

	return lastPrices, nil
}

func (s *service) GetOperations(ctx context.Context, token string, accountID string, from time.Time, to time.Time, instrumentIDs []string) (model.Instruments, model.Operations, error) {
	operations, err := s.tinvestClient.GetOperations(ctx, token, accountID, from, to)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get operations for account %s: %w", accountID, err)
	}

	instruments, err := s.GetInstruments(ctx, token, accountID, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get instruments for account %s: %w", accountID, err)
	}

	for _, operation := range operations {

		instrument, exists := instruments[operation.InstrumentID]
		if !exists {
			continue
		}

		if instrument.OriginalID != "" {
			instrument, exists = instruments[instrument.OriginalID]
			if !exists {
				continue
			}
			operation.InstrumentID = instrument.ID
		}

		// Если валюта RUR
		if operation.Currency == model.CurrencyRUR {
			operation.Currency = model.CurrencyRUB
		}

		// Конвертация цены в рубли
		operation.PriceRub, err = s.exchangeService.Convert(ctx, operation.Currency, model.CurrencyRUB, operation.Price, operation.Time)
		if err != nil {
			return nil, nil, err
		}

		// Конвертация стоимости в рубли
		operation.ValueRub, err = s.exchangeService.Convert(ctx, operation.Currency, model.CurrencyRUB, operation.Value, operation.Time)
		if err != nil {
			return nil, nil, err
		}

		// Конвертация НКД в рубли
		operation.NKDRub, err = s.exchangeService.Convert(ctx, operation.Currency, model.CurrencyRUB, operation.NKD, operation.Time)
		if err != nil {
			return nil, nil, err
		}

		// Конвертация комиссии в рубли
		operation.CommissionRub, err = s.exchangeService.Convert(ctx, operation.Currency, model.CurrencyRUB, operation.Commission, operation.Time)
		if err != nil {
			return nil, nil, err
		}

		// Если валюта операции отличается от валюты инструмента, то
		// сделаем конвертацию суммы операции в валюту инструмента
		if instrument.Currency != operation.Currency {

			operation.Price, err = s.exchangeService.Convert(ctx, operation.Currency, instrument.Currency, operation.Price, operation.Time)
			if err != nil {
				return nil, nil, err
			}

			operation.Value, err = s.exchangeService.Convert(ctx, operation.Currency, instrument.Currency, operation.Value, operation.Time)
			if err != nil {
				return nil, nil, err
			}

			operation.NKD, err = s.exchangeService.Convert(ctx, operation.Currency, instrument.Currency, operation.NKD, operation.Time)
			if err != nil {
				return nil, nil, err
			}

			operation.Commission, err = s.exchangeService.Convert(ctx, operation.Currency, instrument.Currency, operation.Commission, operation.Time)
			if err != nil {
				return nil, nil, err
			}

			operation.Currency = instrument.Currency
		}

		// Payment содержит НКД. Скорректируем payment на сумму НКД
		operation.Value -= operation.NKD
		operation.ValueRub -= operation.NKDRub
	}

	// Сортировка операций по времени
	sort.Slice(operations, func(i, j int) bool {
		return operations[i].Time.Before(operations[j].Time)
	})

	return instruments, operations, nil
}

//nolint:funlen
func (s *service) GetPositions(ctx context.Context, token string, accountID string, from time.Time, to time.Time, instrumentIDs []string) (model.Positions, error) {
	instruments, operations, err := s.GetOperations(ctx, token, accountID, from, to, instrumentIDs)
	if err != nil {
		return nil, err
	}

	positionsMap := map[string]*model.Position{}

	// Операции по бумагам
	for _, operation := range operations {
		if operation.InstrumentID == "" {
			continue
		}

		instrument, exists := instruments[operation.InstrumentID]
		if !exists {
			continue
		}

		if instrument.Type == model.InstrumentTypeCurrency {
			continue
		}

		position, exists := positionsMap[instrument.Ticker]
		if !exists {
			position =
				&model.Position{
					InstrumentID: instrument.ID,
					Ticker:       instrument.Ticker,
					Figi:         instrument.Figi,
					Isin:         instrument.Isin,
					Type:         instrument.Type,
					Name:         instrument.Name,
					Currency:     instrument.Currency,
					Trades:       trades.NewTrades(),
				}

			positionsMap[instrument.Ticker] = position
		}

		switch operation.Type {
		case contractv1.OperationType_OPERATION_TYPE_BUY.String():
			position.QuantityBuy += operation.Quantity
			position.ValueBuy += operation.Value
			position.ValueBuyRub += operation.ValueRub
			position.NKDBuy += operation.NKD
			position.NKDBuyRub += operation.NKDRub
			position.Commissions += operation.Commission
			position.CommissionsRub += operation.CommissionRub

			position.Trades.AddPurchase(operation.Time, operation.Quantity, operation.Value, operation.ValueRub)

		case contractv1.OperationType_OPERATION_TYPE_BUY_CARD.String():
			position.QuantityBuy += operation.Quantity
			position.ValueBuy += operation.Value
			position.ValueBuyRub += operation.ValueRub
			position.NKDBuy += operation.NKD
			position.NKDBuyRub += operation.NKDRub
			position.Commissions += operation.Commission
			position.CommissionsRub += operation.CommissionRub

			position.Trades.AddPurchase(operation.Time, operation.Quantity, operation.Value, operation.ValueRub)

		case contractv1.OperationType_OPERATION_TYPE_SELL.String():
			position.QuantitySell += operation.Quantity
			position.ValueSell += operation.Value
			position.ValueSellRub += operation.ValueRub
			position.NKDSell += operation.NKD
			position.NKDSellRub += operation.NKDRub
			position.Commissions += operation.Commission
			position.CommissionsRub += operation.CommissionRub

			position.Trades.AddSale(operation.Time, operation.Quantity, operation.Value, operation.ValueRub)

		case contractv1.OperationType_OPERATION_TYPE_BROKER_FEE.String():
			// Комиссию берем из операции покупки и продажи
			// position.Commissions += operation.Value
			// position.CommissionsRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_DIVIDEND.String():
			position.Dividends += operation.Value
			position.DividendsRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_DIV_EXT.String():
			position.Dividends += operation.Value
			position.DividendsRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_DIVIDEND_TAX.String():
			position.Taxes += operation.Value
			position.TaxesRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_COUPON.String():
			position.Coupons += operation.Value
			position.CouponsRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_BOND_TAX.String():
			position.Taxes += operation.Value
			position.TaxesRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_TAX.String():
			position.Taxes += operation.Value
			position.TaxesRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_OVERNIGHT.String():
			position.Overnight += operation.Value
			position.OvernightRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_INPUT_SECURITIES.String():
			position.QuantityBuy += operation.Quantity

		case contractv1.OperationType_OPERATION_TYPE_OUT_STAMP_DUTY.String():
			continue

		case contractv1.OperationType_OPERATION_TYPE_INP_MULTI.String():
			continue

		default:
			logger.Warnf(ctx, "Unknown operation type %s: %+v for position %+v", operation.Type, operation, position)
		}

	}

	// Расчет итогов
	for _, position := range positionsMap {
		instrument, exists := instruments[position.InstrumentID]
		if !exists {
			continue
		}
		s.calcPositionTotal(position, instrument)
	}

	// Технические операции
	for _, operation := range operations {
		if operation.InstrumentID != "" {
			continue
		}

		position, exists := positionsMap[""]
		if !exists {
			position =
				&model.Position{
					Currency: operation.Currency,
				}

			positionsMap[""] = position
		}

		switch operation.Type {
		case contractv1.OperationType_OPERATION_TYPE_BOND_TAX.String():
			position.Taxes += operation.Value
			position.TaxesRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_TRACK_MFEE.String():
			position.TrackFee += operation.Value
			position.TrackFeeRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_TRACK_PFEE.String():
			position.ResultFee += operation.Value
			position.ResultFeeRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_OVERNIGHT.String():
			position.Overnight += operation.Value
			position.OvernightRub += operation.ValueRub

		case contractv1.OperationType_OPERATION_TYPE_INPUT.String(), // Пополнение
			contractv1.OperationType_OPERATION_TYPE_OUTPUT.String(), // Вывод
			contractv1.OperationType_OPERATION_TYPE_WRITING_OFF_VARMARGIN.String(),
			contractv1.OperationType_OPERATION_TYPE_ACCRUING_VARMARGIN.String(),
			contractv1.OperationType_OPERATION_TYPE_MARGIN_FEE.String(),
			contractv1.OperationType_OPERATION_TYPE_SERVICE_FEE.String(),
			contractv1.OperationType_OPERATION_TYPE_BROKER_FEE.String(),
			contractv1.OperationType_OPERATION_TYPE_TAX.String(),
			contractv1.OperationType_OPERATION_TYPE_TAX_CORRECTION.String(),
			contractv1.OperationType_OPERATION_TYPE_DIVIDEND_TAX.String(),
			contractv1.OperationType_OPERATION_TYPE_BUY.String(),
			contractv1.OperationType_OPERATION_TYPE_SELL.String(),
			contractv1.OperationType_OPERATION_TYPE_OUT_STAMP_DUTY.String(),
			contractv1.OperationType_OPERATION_TYPE_INP_MULTI.String():
			continue

		default:
			logger.Warnf(ctx, "Unknown operation type %s: %+v\n", operation.Type, *operation)
		}

	}

	// Конвертация ДР
	for _, dr := range s.DrList {
		positionDr, exists := positionsMap[dr.InstrumentTicker]
		if !exists {
			continue
		}

		if positionDr.QuantityEnd > 0 {
			positionDr.QuantitySell += positionDr.QuantityEnd
			positionDr.ValueSell += positionDr.QuantityEnd * positionDr.PriceBuy
			positionDr.ValueSellRub += positionDr.QuantityEnd * positionDr.PriceBuyRub
		}

		position, exists := positionsMap[dr.SourceInstrumentTicker]
		if exists {
			quantity := positionDr.QuantityEnd * dr.Koeff
			value := positionDr.QuantityEnd * positionDr.PriceBuy

			value, err = s.exchangeService.Convert(ctx, model.CurrencyUSD, model.CurrencyRUB, value, s.DrConvTime)
			if err != nil {
				return nil, err
			}

			position.QuantityBuy += quantity
			position.ValueBuy += value
			position.ValueBuyRub += value

			s.calcPositionTotal(position, nil)
		}

		s.calcPositionTotal(positionDr, nil)
	}

	// Подготовка результата
	positions := make(model.Positions, 0, len(positionsMap))
	for _, position := range positionsMap {
		positions = append(positions, position)
	}

	// Сортировка позиций по тикеру
	slices.SortFunc(positions, func(a, b *model.Position) int {
		return strings.Compare(a.Ticker, b.Ticker)
	})

	return positions, nil
}

func (s *service) calcPositionTotal(position *model.Position, instrument *model.Instrument) {
	position.QuantityEnd = position.QuantityBuy - position.QuantitySell

	if position.QuantityBuy != 0 {
		position.PriceBuy = position.ValueBuy / position.QuantityBuy
		position.PriceBuyRub = position.ValueBuyRub / position.QuantityBuy
	}

	if position.QuantitySell != 0 {
		position.PriceSell = position.ValueSell / position.QuantitySell
		position.PriceSellRub = position.ValueSellRub / position.QuantitySell
	}

	if position.QuantityEnd == 0 {
		position.ValueEnd = 0
		position.ValueEndRub = 0
	} else {
		if instrument != nil {
			position.PriceEnd = instrument.LastPrice
			position.PriceEndRub = instrument.LastPriceRub

			position.ValueEnd = position.QuantityEnd * position.PriceEnd
			position.ValueEndRub = position.QuantityEnd * position.PriceEndRub

			position.NKDEnd = position.QuantityEnd * instrument.NKD
			position.NKDEndRub = position.QuantityEnd * instrument.NKDRub
		}
	}

	income := position.ValueEnd +
		position.ValueSell +
		position.NKDSell +
		position.Dividends +
		position.NKDEnd +
		position.Coupons +
		position.Overnight

	incomeRub := position.ValueEndRub +
		position.ValueSellRub +
		position.NKDSellRub +
		position.DividendsRub +
		position.NKDEndRub +
		position.CouponsRub +
		position.OvernightRub

	expenses := position.ValueBuy +
		position.NKDBuy +
		position.Commissions +
		position.Taxes +
		position.TrackFee +
		position.ResultFee

	expensesRub := position.ValueBuyRub +
		position.NKDBuyRub +
		position.CommissionsRub +
		position.TaxesRub +
		position.TrackFeeRub +
		position.ResultFeeRub

	position.Total = income - expenses
	position.TotalRub = incomeRub - expensesRub
}

func (s *service) GetTotals(ctx context.Context, token string, accountID string, from time.Time, to time.Time) (model.Totals, error) {
	positions, err := s.GetPositions(ctx, token, accountID, from, to, nil)
	if err != nil {
		return nil, err
	}

	currencies := map[string]*model.Position{}
	for _, position := range positions {
		currency, exists := currencies[position.Currency]
		if !exists {
			currency = &model.Position{
				Currency: position.Currency,
			}
			currencies[position.Currency] = currency
		}

		currency.ValueBuy += position.ValueBuy
		currency.ValueBuyRub += position.ValueBuyRub
		currency.NKDBuy += position.NKDBuy
		currency.NKDBuyRub += position.NKDBuyRub

		currency.ValueSell += position.ValueSell
		currency.ValueSellRub += position.ValueSellRub
		currency.NKDSell += position.NKDSell
		currency.NKDSellRub += position.NKDSellRub

		currency.ValueEnd += position.ValueEnd
		currency.ValueEndRub += position.ValueEndRub
		currency.NKDEnd += position.NKDEnd
		currency.NKDEndRub += position.NKDEndRub

		currency.Dividends += position.Dividends
		currency.DividendsRub += position.DividendsRub
		currency.Coupons += position.Coupons
		currency.CouponsRub += position.CouponsRub
		currency.Overnight += position.Overnight
		currency.OvernightRub += position.OvernightRub

		currency.Taxes += position.Taxes
		currency.TaxesRub += position.TaxesRub
		currency.Commissions += position.Commissions
		currency.CommissionsRub += position.CommissionsRub
		currency.TrackFee += position.TrackFee
		currency.TrackFeeRub += position.TrackFeeRub
		currency.ResultFee += position.ResultFee
		currency.ResultFeeRub += position.ResultFeeRub
	}

	totals := make(model.Totals, 0, len(currencies))
	for currency, position := range currencies {
		total := &model.Total{
			Currency:       currency,
			ValueBuy:       position.ValueBuy,
			ValueBuyRub:    position.ValueBuyRub,
			NKDBuy:         position.NKDBuy,
			NKDBuyRub:      position.NKDBuyRub,
			ValueSell:      position.ValueSell,
			ValueSellRub:   position.ValueSellRub,
			NKDSell:        position.NKDSell,
			NKDSellRub:     position.NKDSellRub,
			ValueEnd:       position.ValueEnd,
			ValueEndRub:    position.ValueEndRub,
			NKDEnd:         position.NKDEnd,
			NKDEndRub:      position.NKDEndRub,
			Dividends:      position.Dividends,
			DividendsRub:   position.DividendsRub,
			Coupons:        position.Coupons,
			CouponsRub:     position.CouponsRub,
			Overnight:      position.Overnight,
			OvernightRub:   position.OvernightRub,
			Taxes:          position.Taxes,
			TaxesRub:       position.TaxesRub,
			Commissions:    position.Commissions,
			CommissionsRub: position.CommissionsRub,
			TrackFee:       position.TrackFee,
			TrackFeeRub:    position.TrackFeeRub,
			ResultFee:      position.ResultFee,
			ResultFeeRub:   position.ResultFeeRub,
		}

		total.Spent =
			total.ValueBuy +
				total.NKDBuy +
				total.Commissions +
				total.TrackFee +
				total.ResultFee +
				total.Taxes

		total.SpentRub =
			total.ValueBuyRub +
				total.NKDBuyRub +
				total.CommissionsRub +
				total.TrackFeeRub +
				total.ResultFeeRub +
				total.TaxesRub

		total.Received =
			total.ValueSell +
				total.ValueEnd +
				total.NKDSell +
				total.NKDEnd +
				total.Dividends +
				total.Coupons +
				total.Overnight

		total.ReceivedRub =
			total.ValueSellRub +
				total.ValueEndRub +
				total.NKDSellRub +
				total.NKDEndRub +
				total.DividendsRub +
				total.CouponsRub +
				total.OvernightRub

		total.Total = total.Received - total.Spent
		total.TotalRub = total.ReceivedRub - total.SpentRub

		totals = append(totals, total)
	}

	sort.Slice(totals, func(i, j int) bool {
		return cmp.Compare(totals[i].Currency, totals[j].Currency) < 0
	})

	return totals, err

}

func (s *service) GetTrades(ctx context.Context, token string, accountID string, from time.Time, to time.Time) (model.Trades, error) {
	positions, err := s.GetPositions(ctx, token, accountID, from, to, nil)
	if err != nil {
		return nil, err
	}

	var trade *model.Trade
	trades := model.Trades{}
	for _, position := range positions {
		closedTrades := position.Trades.GetClosed()

		for _, closedTrade := range closedTrades {
			if (closedTrade.TimeBuy.After(closedTrade.TimeSell) && from.Before(closedTrade.TimeBuy) && to.After(closedTrade.TimeBuy)) ||
				(closedTrade.TimeSell.After(closedTrade.TimeBuy) && from.Before(closedTrade.TimeSell) && to.After(closedTrade.TimeSell)) {

				trade = &model.Trade{
					Ticker:   position.Ticker,
					Currency: position.Currency,
					Trade:    closedTrade,
				}

				if closedTrade.TimeBuy.Before(closedTrade.TimeSell) {
					trade.Time = closedTrade.TimeSell
				} else {
					trade.Time = closedTrade.TimeBuy
				}

				trades = append(trades, trade)
			}
		}
	}

	// Сортировка времени
	sort.Slice(trades, func(i, j int) bool {
		return trades[i].Time.Before(trades[j].Time)
	})

	return trades, nil
}

func (s *service) GetCandles(ctx context.Context, token string, instrumentID string, interval contractv1.CandleInterval, from time.Time, to time.Time) (model.Candles, error) {
	candles, err := s.tinvestClient.GetCandles(ctx, token, instrumentID, interval, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get candles: %w", err)
	}
	return candles, nil
}
