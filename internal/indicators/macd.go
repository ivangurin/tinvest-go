package indicators

import (
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"

	"tinvest-go/internal/model"
)

func GetMACD(candles model.Candles) MACDResult {
	techanTimeSeries := techan.NewTimeSeries()
	for _, candle := range candles {
		techanCandle := techan.NewCandle(techan.TimePeriod{Start: candle.Time, End: candle.Time})

		techanCandle.OpenPrice = big.NewDecimal(candle.Open)
		techanCandle.MinPrice = big.NewDecimal(candle.Low)
		techanCandle.MaxPrice = big.NewDecimal(candle.High)
		techanCandle.ClosePrice = big.NewDecimal(candle.Close)
		techanCandle.Volume = big.NewDecimal(float64(candle.Volume))

		techanTimeSeries.AddCandle(techanCandle)
	}

	techanClosePriceIndicator := techan.NewClosePriceIndicator(techanTimeSeries)
	techanMACDIndicator := techan.NewMACDIndicator(techanClosePriceIndicator, 12, 26)
	techanMACDHistogramIndicator := techan.NewMACDHistogramIndicator(techanMACDIndicator, 9)

	macdResult := make(MACDResult, 0, len(candles))
	for index, candle := range candles {
		macdValue :=
			&MACDValue{
				Time:  candle.Time,
				Value: techanMACDHistogramIndicator.Calculate(index).Float(),
			}

		macdResult = append(macdResult, macdValue)
	}

	return macdResult
}
