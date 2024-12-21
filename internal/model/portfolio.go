package model

type PortfolioPosition struct {
	InstrumentID string
	Isin         string
	Ticker       string
	Name         string
	Quantity     float64
	Currency     string
	Price        float64
	PriceRub     float64
	Value        float64
	ValueRub     float64
	PriceEnd     float64
	PriceEndRub  float64
	ValueEnd     float64
	ValueEndRub  float64
	Total        float64
	TotalRub     float64
	Percent      float64
	PercentRub   float64
}

type Portfolio []*PortfolioPosition
