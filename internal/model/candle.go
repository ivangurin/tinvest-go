package model

import "time"

type Candle struct {
	Time   time.Time
	Open   float64
	Low    float64
	High   float64
	Close  float64
	Volume int64
}

type Candles []*Candle
