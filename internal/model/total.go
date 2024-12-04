package model

type Total struct {
	Currency       string
	ValueBuy       float64
	ValueBuyRub    float64
	NKDBuy         float64
	NKDBuyRub      float64
	ValueSell      float64
	ValueSellRub   float64
	NKDSell        float64
	NKDSellRub     float64
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
}

type Totals []*Total
