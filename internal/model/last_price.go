package model

type LastPrice struct {
	Value         float64
	AbsoluteValue bool
}

type LastPrices map[string]*LastPrice
