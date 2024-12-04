package model

import "tinvest-go/internal/pkg/trades"

type Position struct {
	InstrumentID   string
	Ticker         string
	Figi           string
	Isin           string
	Type           string
	Name           string
	Currency       string
	QuantityBuy    float64
	PriceBuy       float64
	PriceBuyRub    float64
	ValueBuy       float64
	ValueBuyRub    float64
	NKDBuy         float64
	NKDBuyRub      float64
	QuantitySell   float64
	PriceSell      float64
	PriceSellRub   float64
	ValueSell      float64
	ValueSellRub   float64
	NKDSell        float64
	NKDSellRub     float64
	QuantityEnd    float64
	PriceEnd       float64
	PriceEndRub    float64
	ValueEnd       float64
	ValueEndRub    float64
	NKDEnd         float64
	NKDEndRub      float64
	Dividends      float64
	DividendsRub   float64
	Coupons        float64
	CouponsRub     float64
	Overnight      float64
	OvernightRub   float64
	Taxes          float64
	TaxesRub       float64
	Commissions    float64
	CommissionsRub float64
	TrackFee       float64
	TrackFeeRub    float64
	ResultFee      float64
	ResultFeeRub   float64
	Spent          float64
	SpentRub       float64
	Received       float64
	ReceivedRub    float64
	Total          float64
	TotalRub       float64
	Trades         trades.Trades
}

type Positions []*Position
