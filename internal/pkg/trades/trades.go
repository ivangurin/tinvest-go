package trades

import "time"

type Trades struct {
	trades TradeList
}

type Trade struct {
	TimeBuy      time.Time
	QuantityBuy  float64
	PriceBuy     float64
	PriceBuyRub  float64
	ValueBuy     float64
	ValueBuyRub  float64
	TimeSell     time.Time
	QuantitySell float64
	PriceSell    float64
	PriceSellRub float64
	ValueSell    float64
	ValueSellRub float64
	Total        float64
	TotalRub     float64
	Percent      float64
	PercentRub   float64
}

type TradeList []*Trade

func NewTrades() Trades {

	return Trades{}

}

func (l *Trades) GetAll() TradeList {
	return l.trades
}

func (l *Trades) GetClosed() TradeList {

	trades := TradeList{}

	for _, trade := range l.trades {
		if trade.QuantityBuy != 0 &&
			trade.QuantitySell != 0 {
			trades = append(trades, trade)
		}
	}

	return trades

}

func (l *Trades) GetUnclosed() TradeList {

	trades := TradeList{}

	for _, trade := range l.trades {
		if trade.Total == 0 {
			trades = append(trades, trade)
		}
	}

	return trades

}

func (l *Trades) AddPurchase(time time.Time, quantity float64, value float64, valueRub float64) {
	if quantity == 0 {
		return
	}

	price := value / quantity
	priceRub := valueRub / quantity

	// Будем выполянть столько итераций пока не распределеим все количество
	for {

		// Для каждой сделки
		for index, trade := range l.trades {

			// Если колчиество купленного равно 0,
			// то это занчит, что сделка завершена и ее пропускаем
			if trade.QuantityBuy != 0 {
				continue
			}

			// Если количество равно проданному количеству, то
			// заполняем полностью купленное количество
			if quantity == trade.QuantitySell {

				trade.TimeBuy = time
				trade.QuantityBuy = quantity
				trade.PriceBuy = price
				trade.PriceBuyRub = priceRub
				trade.ValueBuy = trade.QuantityBuy * price
				trade.ValueBuyRub = trade.QuantityBuy * priceRub
				trade.Total = trade.ValueSell - trade.ValueBuy
				trade.TotalRub = trade.ValueSellRub - trade.ValueBuyRub
				trade.Percent = trade.Total / trade.ValueSell * 100
				trade.PercentRub = trade.TotalRub / trade.ValueSellRub * 100

				quantity = 0

				// Если количество больше чем проданное количество, то
				// заполним купленное количество проданным количеством.
				// Остаток количества распределится на другие сделки
			} else if quantity > trade.QuantitySell {

				trade.TimeBuy = time
				trade.QuantityBuy = trade.QuantitySell
				trade.PriceBuy = price
				trade.PriceBuyRub = priceRub
				trade.ValueBuy = trade.QuantityBuy * price
				trade.ValueBuyRub = trade.QuantityBuy * priceRub
				trade.Total = trade.ValueSell - trade.ValueBuy
				trade.TotalRub = trade.ValueSellRub - trade.ValueBuyRub
				trade.Percent = trade.Total / trade.ValueSell * 100
				trade.PercentRub = trade.TotalRub / trade.ValueSellRub * 100

				quantity -= trade.QuantityBuy

				// Если количество меньше чем проданное количество, то
				// заполним купленное количество всем количеством
				// И для остатка проданнго количества добавим новую сделку
			} else if quantity < trade.QuantitySell {

				trade.TimeBuy = time
				trade.QuantityBuy = quantity
				trade.PriceBuy = price
				trade.PriceBuyRub = priceRub
				trade.ValueBuy = quantity * price
				trade.ValueBuyRub = quantity * priceRub

				newTrade := &Trade{
					TimeSell:     trade.TimeSell,
					QuantitySell: trade.QuantitySell - quantity,
					PriceSell:    trade.PriceSell,
					PriceSellRub: trade.PriceSellRub,
					ValueSell:    (trade.QuantitySell - quantity) * trade.PriceSell,
					ValueSellRub: (trade.QuantitySell - quantity) * trade.PriceSellRub}

				trade.QuantitySell = quantity
				trade.ValueSell = quantity * trade.PriceSell
				trade.ValueSellRub = quantity * trade.PriceSellRub
				trade.Total = trade.ValueSell - trade.ValueBuy
				trade.TotalRub = trade.ValueSell - trade.ValueBuyRub
				trade.Percent = trade.Total / trade.ValueSell * 100
				trade.PercentRub = trade.TotalRub / trade.ValueSellRub * 100

				quantity = 0

				tradesTmp := append(l.trades[:index+1], l.trades[index:]...)
				tradesTmp[index+1] = newTrade

				l.trades = tradesTmp

			}

			if quantity == 0 {
				break
			}

		}

		// Если полсе распределения осталось количество, то добавим его как нувую сделку
		if quantity != 0 {

			trade := &Trade{
				TimeBuy:     time,
				QuantityBuy: quantity,
				PriceBuy:    price,
				PriceBuyRub: priceRub,
				ValueBuy:    quantity * price,
				ValueBuyRub: quantity * priceRub,
			}

			l.trades = append(l.trades, trade)

			quantity = 0

		}

		if quantity == 0 {
			break
		}

	}

}

func (l *Trades) AddSale(date time.Time, quantity float64, value float64, valueRub float64) {

	if quantity == 0 {
		return
	}

	price := value / quantity
	priceRub := valueRub / quantity

	// Будем выполянть столько итераций пока не распределеим все количество
	for {

		// Для каждой сделки
		for index, trade := range l.trades {

			// Если колчиество проданног не равно 0,
			// то это занчит, что сделка завершена и ее пропускаем
			if trade.QuantitySell != 0 {
				continue
			}

			// Если количество равно купленному количеству, то
			// заполняем полностью проданное количество
			if quantity == trade.QuantityBuy {

				trade.TimeSell = date
				trade.QuantitySell = quantity
				trade.PriceSell = price
				trade.PriceSellRub = priceRub
				trade.ValueSell = trade.QuantitySell * price
				trade.ValueSellRub = trade.QuantitySell * priceRub
				trade.Total = trade.ValueSell - trade.ValueBuy
				trade.TotalRub = trade.ValueSellRub - trade.ValueBuyRub
				trade.Percent = trade.Total / trade.ValueBuy * 100
				trade.PercentRub = trade.TotalRub / trade.ValueBuyRub * 100

				quantity = 0

				// Если количество больше чем купленное количество, то
				// заполним проданное количество купленным количеством.
				// Остаток количества распределится на другие сделки
			} else if quantity > trade.QuantityBuy {

				trade.TimeSell = date
				trade.QuantitySell = trade.QuantityBuy
				trade.PriceSell = price
				trade.PriceSellRub = priceRub
				trade.ValueSell = trade.QuantitySell * price
				trade.ValueSellRub = trade.QuantitySell * priceRub
				trade.Total = trade.ValueSell - trade.ValueBuy
				trade.TotalRub = trade.ValueSellRub - trade.ValueBuyRub
				trade.Percent = trade.Total / trade.ValueBuy * 100
				trade.PercentRub = trade.TotalRub / trade.ValueBuyRub * 100

				quantity -= trade.QuantitySell

				// Если количество меньше чем купленное количество, то
				// заполним проданное количество всем количеством
				// И для остатка купленного количества добавим новую сделку
			} else if quantity < trade.QuantityBuy {

				trade.TimeSell = date
				trade.QuantitySell = quantity
				trade.PriceSell = price
				trade.PriceSellRub = priceRub
				trade.ValueSell = quantity * price
				trade.ValueSellRub = quantity * priceRub

				newTrade := &Trade{
					TimeBuy:     trade.TimeBuy,
					QuantityBuy: trade.QuantityBuy - quantity,
					PriceBuy:    trade.PriceBuy,
					PriceBuyRub: trade.PriceBuyRub,
					ValueBuy:    (trade.QuantityBuy - quantity) * trade.PriceBuy,
					ValueBuyRub: (trade.QuantityBuy - quantity) * trade.PriceBuyRub}

				trade.QuantityBuy = quantity
				trade.ValueBuy = quantity * trade.PriceBuy
				trade.ValueBuyRub = quantity * trade.PriceBuyRub
				trade.Total = trade.ValueSell - trade.ValueBuy
				trade.TotalRub = trade.ValueSellRub - trade.ValueBuyRub
				trade.Percent = trade.Total / trade.ValueBuy * 100
				trade.PercentRub = trade.TotalRub / trade.ValueBuyRub * 100

				quantity = 0

				tradesTmp := append(l.trades[:index+1], l.trades[index:]...)
				tradesTmp[index+1] = newTrade

				l.trades = tradesTmp

			}

			if quantity == 0 {
				break
			}

		}

		// Если полсе распределения осталось количество, то добавим его как нувую сделку
		if quantity != 0 {

			trade := &Trade{
				TimeSell:     date,
				QuantitySell: quantity,
				PriceSell:    price,
				PriceSellRub: priceRub,
				ValueSell:    quantity * price,
				ValueSellRub: quantity * priceRub,
			}

			l.trades = append(l.trades, trade)

			quantity = 0

		}

		if quantity == 0 {
			break
		}

	}

}
