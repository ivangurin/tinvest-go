package model

import (
	"time"
	"tinvest-go/internal/pkg/trades"
)

type Trade struct {
	Time     time.Time
	Ticker   string
	Currency string
	Trade    *trades.Trade
}

type Trades []*Trade
