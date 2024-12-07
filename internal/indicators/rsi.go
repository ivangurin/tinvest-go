package indicators

import (
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"

	"tinvest-go/internal/model"
)

func GetRSI(candles model.Candles) RSIResult {
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
	techanRSIIndicator := techan.NewRelativeStrengthIndexIndicator(techanClosePriceIndicator, 14)

	rsiResult := make(RSIResult, 0, len(candles))
	for index, candle := range candles {
		rsiValue :=
			&RSIValue{
				Time:  candle.Time,
				Value: techanRSIIndicator.Calculate(index).Float(),
			}

		rsiResult = append(rsiResult, rsiValue)
	}

	return rsiResult
}
