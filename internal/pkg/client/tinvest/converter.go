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

func convertInstrumentType(instrumentType contractv1.InstrumentType) string {
	var res string
	switch instrumentType {
	case contractv1.InstrumentType_INSTRUMENT_TYPE_CURRENCY:
		res = model.InstrumentTypeCurrency
	case contractv1.InstrumentType_INSTRUMENT_TYPE_SHARE:
		res = model.InstrumentTypeShare
	case contractv1.InstrumentType_INSTRUMENT_TYPE_BOND:
		res = model.InstrumentTypeBond
	case contractv1.InstrumentType_INSTRUMENT_TYPE_ETF:
		res = model.InstrumentTypeEtf
	case contractv1.InstrumentType_INSTRUMENT_TYPE_FUTURES:
		res = model.InstrumentTypeFuture
	default:
		res = instrumentType.String()
	}
	return res
}

func convertInstrument(instrument *contractv1.Instrument) *model.Instrument {
	return &model.Instrument{
		ID:       instrument.GetUid(),
		Ticker:   instrument.GetTicker(),
		Figi:     instrument.GetFigi(),
		Isin:     instrument.GetIsin(),
		Name:     instrument.GetName(),
		Type:     convertInstrumentType(instrument.GetInstrumentKind()),
		Currency: strings.ToUpper(instrument.GetCurrency()),
		Lot:      instrument.GetLot(),
		Country:  instrument.GetCountryOfRisk(),
		Trading:  instrument.GetApiTradeAvailableFlag(),
	}
}

func convertLastPrice(lastPrice *contractv1.LastPrice) *model.LastPrice {
	return &model.LastPrice{
		Value: utils.QuotationToFloat64(lastPrice.GetPrice()),
	}
}

func convertCurrencies(currencies []*contractv1.Currency) model.Instruments {
	res := make(model.Instruments, 0, len(currencies))
	for _, currency := range currencies {
		res = append(res, convertCurrency(currency))
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
		Type:     model.InstrumentTypeCurrency,
		Currency: strings.ToUpper(currency.GetCurrency()),
		Lot:      currency.GetLot(),
		Country:  currency.GetCountryOfRisk(),
		Trading:  currency.GetApiTradeAvailableFlag(),
		Nominal:  utils.MoneyToFloat64(currency.GetNominal()),
	}
}

func convertShares(currencies []*contractv1.Share) model.Instruments {
	res := make(model.Instruments, 0, len(currencies))
	for _, share := range currencies {
		res = append(res, convertShare(share))
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
		Type:     model.InstrumentTypeShare,
		Currency: strings.ToUpper(share.GetCurrency()),
		Lot:      share.GetLot(),
		Country:  share.GetCountryOfRisk(),
		Trading:  share.GetApiTradeAvailableFlag(),
	}
}

func convertBonds(bonds []*contractv1.Bond) model.Instruments {
	res := make(model.Instruments, 0, len(bonds))
	for _, bond := range bonds {
		res = append(res, convertBond(bond))
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
		Type:     model.InstrumentTypeBond,
		Currency: strings.ToUpper(bond.GetCurrency()),
		Lot:      bond.GetLot(),
		Country:  bond.GetCountryOfRisk(),
		Trading:  bond.GetApiTradeAvailableFlag(),
		Nominal:  utils.MoneyToFloat64(bond.GetNominal()),
		NKD:      utils.MoneyToFloat64(bond.GetAciValue()),
	}
}

func convertEtfs(etfs []*contractv1.Etf) model.Instruments {
	res := make(model.Instruments, 0, len(etfs))
	for _, etf := range etfs {
		res = append(res, convertEtf(etf))
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
		Type:     model.InstrumentTypeEtf,
		Currency: strings.ToUpper(etf.GetCurrency()),
		Lot:      etf.GetLot(),
		Country:  etf.GetCountryOfRisk(),
		Trading:  etf.GetApiTradeAvailableFlag(),
	}
}

func convertFutures(futures []*contractv1.Future) model.Instruments {
	res := make(model.Instruments, 0, len(futures))
	for _, future := range futures {
		res = append(res, convertFuture(future))
	}
	return res
}

func convertFuture(future *contractv1.Future) *model.Instrument {
	return &model.Instrument{
		ID:                      future.GetUid(),
		Ticker:                  future.GetTicker(),
		Figi:                    future.GetFigi(),
		Name:                    future.GetName(),
		Type:                    model.InstrumentTypeFuture,
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

func convertOperation(operation *contractv1.OperationItem) *model.Operation {
	return &model.Operation{
		ID:           operation.GetId(),
		Time:         operation.GetDate().AsTime(),
		Name:         operation.GetName(),
		Type:         operation.GetType().String(),
		InstrumentID: operation.GetInstrumentUid(),
		PositionID:   operation.GetPositionUid(),
		Figi:         operation.GetFigi(),
		Quantity:     float64(operation.GetQuantityDone()),
		Price:        math.Abs(utils.MoneyToFloat64(operation.Price)),
		Value:        math.Abs(utils.MoneyToFloat64(operation.Payment)),
		NKD:          math.Abs(utils.MoneyToFloat64(operation.AccruedInt)),
		Commission:   math.Abs(utils.MoneyToFloat64(operation.Commission)),
		Currency:     strings.ToUpper(operation.GetPayment().GetCurrency()),
	}
}
