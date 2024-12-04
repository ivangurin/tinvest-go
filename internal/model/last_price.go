package model

type LastPrice struct {
	Value    float64
	Absolute bool
}

type LastPrices map[string]*LastPrice
