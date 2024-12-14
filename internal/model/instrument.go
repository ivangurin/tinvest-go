package model

const (
	InstrumentTypeCurrency = "currency"
	InstrumentTypeShare    = "share"
	InstrumentTypeBond     = "bond"
	InstrumentTypeEtf      = "etf"
	InstrumentTypeFuture   = "future"
	InstrumentTypeOption   = "option"
)

type Instrument struct {
	ID                      string
	Type                    string
	Ticker                  string
	Name                    string
	Figi                    string
	FigiOrig                string
	Isin                    string
	Currency                string
	Lot                     int32
	Country                 string
	Trading                 bool
	NKD                     float64
	NKDRub                  float64
	LastPrice               float64
	LastPriceRub            float64
	Nominal                 float64
	MinPriceIncrement       float64
	MinPriceIncrementAmount float64
}

type Instruments map[string]*Instrument
