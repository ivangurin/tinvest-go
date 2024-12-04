package tinvest_service

type DrItem struct {
	InstrumentID       string
	SourceInstrumentID string
	Koeff              float64
}

type DrList []*DrItem
