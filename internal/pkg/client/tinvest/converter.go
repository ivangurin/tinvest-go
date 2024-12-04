package tinvest_client

import (
	"cmp"
	"math"
	"slices"
	"strings"
	"tinvest-go/internal/model"
	contractv1 "tinvest-go/internal/pb"
	"tinvest-go/internal/pkg/utils"
)

func convertAccounts(accounts []*contractv1.Account) model.Accounts {
	res := make(model.Accounts, 0, len(accounts))
	for _, account := range accounts {
		res = append(res, convertAccount(account))
	}
	slices.SortFunc(res, func(a, b *model.Account) int {
		return cmp.Compare(a.ID, b.ID)
	})
	return res
}

func convertAccount(account *contractv1.Account) *model.Account {
	return &model.Account{
		ID:       account.GetId(),
		Type:     account.GetType(),
		Name:     account.GetName(),
		OpenedAt: account.GetOpenedDate().AsTime(),
		ClosedAt: utils.IfThenElse(account.GetClosedDate().AsTime().After(account.GetOpenedDate().AsTime()), utils.Ptr(account.GetClosedDate().AsTime()), nil),
	}
}

func convertInstrument(instrument *contractv1.Instrument) *model.Instrument {
	return &model.Instrument{
		ID:       instrument.GetUid(),
		Ticker:   instrument.GetTicker(),
		Figi:     instrument.GetFigi(),
		Isin:     instrument.GetIsin(),
		Name:     instrument.GetName(),
		Type:     instrument.GetInstrumentKind().String(),
		Currency: strings.ToUpper(instrument.GetCurrency()),
		Lot:      instrument.GetLot(),
		Country:  instrument.GetCountryOfRisk(),
		Trading:  instrument.GetApiTradeAvailableFlag(),
	}
}

func convertLastPrices(lastPrices []*contractv1.LastPrice) model.LastPrices {
	res := make(model.LastPrices, len(lastPrices))
	for _, lastPrice := range lastPrices {
		res[lastPrice.GetInstrumentUid()] = convertLastPrice(lastPrice)
	}
	return res
}

func convertLastPrice(lastPrice *contractv1.LastPrice) *model.LastPrice {
	return &model.LastPrice{
		Value: utils.QuotationToFloat64(lastPrice.GetPrice()),
	}
}

func convertCurrencies(currencies []*contractv1.Currency) model.Instruments {
	res := make(model.Instruments, len(currencies))
	for _, currency := range currencies {
		res[currency.GetUid()] = convertCurrency(currency)
	}
	return res
}

func convertCurrency(currency *contractv1.Currency) *model.Instrument {
	return &model.Instrument{
		ID:       currency.GetUid(),
		Ticker:   currency.GetTicker(),
		Figi:     currency.GetFigi(),
		Isin:     currency.GetIsin(),
		Name:     currency.GetName(),
		Type:     contractv1.InstrumentType_INSTRUMENT_TYPE_CURRENCY.String(),
		Currency: strings.ToUpper(currency.GetCurrency()),
		Lot:      currency.GetLot(),
		Country:  currency.GetCountryOfRisk(),
		Trading:  currency.GetApiTradeAvailableFlag(),
		Nominal:  utils.MoneyToFloat64(currency.GetNominal()),
	}
}

func convertShares(currencies []*contractv1.Share) model.Instruments {
	res := make(model.Instruments, len(currencies))
	for _, share := range currencies {
		res[share.GetUid()] = convertShare(share)
	}
	return res
}

func convertShare(share *contractv1.Share) *model.Instrument {
	return &model.Instrument{
		ID:       share.GetUid(),
		Ticker:   share.GetTicker(),
		Figi:     share.GetFigi(),
		Isin:     share.GetIsin(),
		Name:     share.GetName(),
		Type:     contractv1.InstrumentType_INSTRUMENT_TYPE_SHARE.String(),
		Currency: strings.ToUpper(share.GetCurrency()),
		Lot:      share.GetLot(),
		Country:  share.GetCountryOfRisk(),
		Trading:  share.GetApiTradeAvailableFlag(),
	}
}

func convertBonds(bonds []*contractv1.Bond) model.Instruments {
	res := make(model.Instruments, len(bonds))
	for _, bond := range bonds {
		res[bond.GetUid()] = convertBond(bond)
	}
	return res
}

func convertBond(bond *contractv1.Bond) *model.Instrument {
	return &model.Instrument{
		ID:       bond.GetUid(),
		Ticker:   bond.GetTicker(),
		Figi:     bond.GetFigi(),
		Isin:     bond.GetIsin(),
		Name:     bond.GetName(),
		Type:     contractv1.InstrumentType_INSTRUMENT_TYPE_BOND.String(),
		Currency: strings.ToUpper(bond.GetCurrency()),
		Lot:      bond.GetLot(),
		Country:  bond.GetCountryOfRisk(),
		Trading:  bond.GetApiTradeAvailableFlag(),
		Nominal:  utils.MoneyToFloat64(bond.GetNominal()),
		NKD:      utils.MoneyToFloat64(bond.GetAciValue()),
	}
}

func convertEtfs(etfs []*contractv1.Etf) model.Instruments {
	res := make(model.Instruments, len(etfs))
	for _, etf := range etfs {
		res[etf.GetUid()] = convertEtf(etf)
	}
	return res
}

func convertEtf(etf *contractv1.Etf) *model.Instrument {
	return &model.Instrument{
		ID:       etf.GetUid(),
		Ticker:   etf.GetTicker(),
		Figi:     etf.GetFigi(),
		Isin:     etf.GetIsin(),
		Name:     etf.GetName(),
		Type:     contractv1.InstrumentType_INSTRUMENT_TYPE_ETF.String(),
		Currency: strings.ToUpper(etf.GetCurrency()),
		Lot:      etf.GetLot(),
		Country:  etf.GetCountryOfRisk(),
		Trading:  etf.GetApiTradeAvailableFlag(),
	}
}

func convertFutures(futures []*contractv1.Future) model.Instruments {
	res := make(model.Instruments, len(futures))
	for _, future := range futures {
		res[future.GetUid()] = convertFuture(future)
	}
	return res
}

func convertFuture(future *contractv1.Future) *model.Instrument {
	return &model.Instrument{
		ID:                      future.GetUid(),
		Ticker:                  future.GetTicker(),
		Figi:                    future.GetFigi(),
		Name:                    future.GetName(),
		Type:                    contractv1.InstrumentType_INSTRUMENT_TYPE_FUTURES.String(),
		Currency:                strings.ToUpper(future.GetCurrency()),
		Lot:                     future.GetLot(),
		Country:                 future.GetCountryOfRisk(),
		Trading:                 future.GetApiTradeAvailableFlag(),
		MinPriceIncrement:       utils.QuotationToFloat64(future.GetMinPriceIncrement()),
		MinPriceIncrementAmount: utils.QuotationToFloat64(future.GetMinPriceIncrementAmount()),
	}
}

func convertFavorites(favorites []*contractv1.FavoriteInstrument) model.Favorites {
	res := make(model.Favorites, 0, len(favorites))
	for _, favorite := range favorites {
		res = append(res, convertFavorite(favorite))
	}
	return res
}

func convertFavorite(favorite *contractv1.FavoriteInstrument) *model.Favorite {
	return &model.Favorite{
		ID:     favorite.GetUid(),
		Kind:   favorite.GetInstrumentKind().String(),
		Ticker: favorite.GetTicker(),
	}
}

func convertPortfolioPositions(portfolioPositions []*contractv1.PortfolioPosition) PortfolioPositions {
	res := make(PortfolioPositions, 0, len(portfolioPositions))
	for _, portfolioPosition := range portfolioPositions {
		res = append(res, convertPortfolioPosition(portfolioPosition))
	}
	return res
}

func convertPortfolioPosition(portfolioPosition *contractv1.PortfolioPosition) *PortfolioPosition {
	return &PortfolioPosition{
		ID:           portfolioPosition.GetInstrumentUid(),
		Type:         strings.ToUpper(portfolioPosition.GetInstrumentType()),
		CurrentPrice: utils.MoneyToFloat64(portfolioPosition.GetCurrentPrice()),
		CurrencyCode: strings.ToUpper(portfolioPosition.GetCurrentPrice().GetCurrency()),
	}
}

func convertOperations(operations []*contractv1.OperationItem) model.Operations {
	res := make(model.Operations, 0, len(operations))
	for _, operation := range operations {
		res = append(res, convertOperation(operation))
	}
	return res
}

func convertOperation(operation *contractv1.OperationItem) *model.Operation {
	return &model.Operation{
		ID:           operation.GetId(),
		Type:         operation.GetType().String(),
		Time:         operation.GetDate().AsTime(),
		InstrumentID: operation.GetInstrumentUid(),
		Quantity:     float64(operation.GetQuantityDone()),
		Price:        math.Abs(utils.MoneyToFloat64(operation.Price)),
		Value:        math.Abs(utils.MoneyToFloat64(operation.Payment)),
		NKD:          math.Abs(utils.MoneyToFloat64(operation.AccruedInt)),
		Commission:   math.Abs(utils.MoneyToFloat64(operation.Commission)),
		Currency:     strings.ToUpper(operation.GetPayment().GetCurrency()),
	}
}
