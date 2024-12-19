package model

import (
	"time"
	"tinvest-go/internal/pkg/utils"
)

type Operation struct {
	ID            string
	Time          time.Time
	Name          string
	Type          string
	InstrumentID  string
	PositionID    string
	Figi          string
	Quantity      float64
	Price         float64
	PriceRub      float64
	Value         float64
	ValueRub      float64
	NKD           float64
	NKDRub        float64
	Commission    float64
	CommissionRub float64
	Currency      string
}

type Operations []*Operation

func (o *Operations) GetInstrumentsIDs() []string {
	set := utils.ToSet(*o, func(operation *Operation) string {
		return operation.InstrumentID
	})
	return set.ToSlice()
}
