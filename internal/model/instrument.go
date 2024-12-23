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
	OriginalID              string
	Isin                    string
	Figi                    string
	Ticker                  string
	Type                    string
	Name                    string
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

type Instruments []*Instrument

func (i *Instruments) GetIDs() []string {
	res := make([]string, 0, len(*i))
	for _, instrument := range *i {
		res = append(res, instrument.ID)
	}
	return res
}
