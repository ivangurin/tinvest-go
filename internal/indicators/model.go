package indicators

import "time"

type RSIValue struct {
	Time  time.Time
	Value float64
}

type RSIResult []*RSIValue

type MACDValue struct {
	Time  time.Time
	Value float64
}

type MACDResult []*MACDValue
