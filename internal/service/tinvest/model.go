package tinvest_service

type DrItem struct {
	InstrumentTicker       string
	SourceInstrumentTicker string
	Koeff                  float64
}

type DrList []*DrItem
