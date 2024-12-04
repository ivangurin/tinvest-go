package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountTradesFor(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
	// exists, err := checkToken(user)
	// if err != nil {
	// 	return err
	// }

	// if !exists {
	// 	return nil
	// }

	// message := tgbotapi.NewMessage(user.ChatID, texts.Processing)
	// messageID, err := a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	// defer deleteMessage(user.ChatID, messageID)

	// now := time.Now()

	// weekday := now.Weekday()

	// if weekday == 0 {
	// 	weekday = 7
	// }

	// var title string
	// var from time.Time
	// var to time.Time

	// switch period {
	// case "currDay":

	// 	title = texts.AccountTradesCurrDay

	// 	from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	// 	to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	// case "prevDay":

	// 	title = texts.AccountTradesPrevDay

	// 	from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	// 	from = from.AddDate(0, 0, -1)

	// 	to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	// 	to = to.AddDate(0, 0, -1)

	// case "currWeek":

	// 	title = texts.AccountTradesCurrWeek

	// 	from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	// 	from = from.AddDate(0, 0, -int(weekday-1))

	// 	to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	// case "prevWeek":

	// 	title = texts.AccountTradesPrevWeek

	// 	from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	// 	from = from.AddDate(0, 0, -int(7+weekday-1))

	// 	to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	// 	to = to.AddDate(0, 0, -int(weekday))

	// case "currMonth":

	// 	title = texts.AccountTradesCurrMonth

	// 	from = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	// 	to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	// case "prevMonth":

	// 	title = texts.AccountTradesPrevMonth

	// 	to = time.Date(now.Year(), now.Month(), 1, 23, 59, 59, 0, time.UTC)
	// 	to = to.AddDate(0, 0, -1)

	// 	from = time.Date(to.Year(), to.Month(), 1, 0, 0, 0, 0, time.UTC)

	// default:
	// 	return nil
	// }

	// tiClient, err := tinvest.NewClient(user.Token, repo)
	// if err != nil {
	// 	return err
	// }

	// account, err := tiClient.GetAccount(accountId)
	// if err != nil {
	// 	return err
	// }

	// if account == nil {
	// 	return nil
	// }

	// tiClient.SetAccountId(account.Id)

	// trades, err := tiClient.GetTrades(from, to)
	// if err != nil {
	// 	return err
	// }

	// sb := strings.Builder{}

	// sb.WriteString(fmt.Sprintf(texts.AccountTradesTitle, title, from.Format("02.01.2006"), to.Format("02.01.2006"), account.Name))
	// sb.WriteString("\n")

	// if len(trades) == 0 {

	// 	sb.WriteString(texts.AccountTradesNoTrades)

	// 	message = tgbotapi.NewMessage(user.ChatID, "")
	// 	message.Text = sb.String()

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil

	// }

	// group := map[time.Time]tinvest.Trades{}

	// for _, trade := range trades {

	// 	date := time.Date(trade.Time.Year(), trade.Time.Month(), trade.Time.Day(), 0, 0, 0, 0, time.UTC)

	// 	group[date] = append(group[date], trade)

	// }

	// keys := []time.Time{}

	// for key := range group {
	// 	keys = append(keys, key)
	// }

	// sort.Slice(keys, func(i, j int) bool {
	// 	return keys[i].Before(keys[j])
	// })

	// totalRub := float64(0)

	// for keyIndex, key := range keys {

	// 	if keyIndex > 0 {
	// 		sb = strings.Builder{}
	// 	}

	// 	sb.WriteString(fmt.Sprintf("<b>%s:</b>\n", key.Format("02.01.2006")))

	// 	trades = group[key]

	// 	totalDayRub := float64(0)

	// 	for _, trade := range trades {

	// 		totalRub += trade.Trade.TotalRub
	// 		totalDayRub += trade.Trade.TotalRub

	// 		if trade.Currency == tinvest.CurrencyRUB {

	// 			sb.WriteString(fmt.Sprintf(texts.AccountTradesItemRub,
	// 				trade.Ticker,
	// 				trade.Trade.QuantityBuy,
	// 				trade.Currency,
	// 				ac.FormatMoney(trade.Trade.ValueBuy),
	// 				ac.FormatMoney(trade.Trade.ValueSell),
	// 				ac.FormatMoney(trade.Trade.Total),
	// 				trade.Trade.Percent))

	// 		} else {

	// 			sb.WriteString(fmt.Sprintf(texts.AccountTradesItem,
	// 				trade.Ticker,
	// 				trade.Trade.QuantityBuy,
	// 				trade.Currency,
	// 				ac.FormatMoney(trade.Trade.ValueBuy),
	// 				ac.FormatMoney(trade.Trade.ValueBuyRub),
	// 				ac.FormatMoney(trade.Trade.ValueSell),
	// 				ac.FormatMoney(trade.Trade.ValueSellRub),
	// 				ac.FormatMoney(trade.Trade.Total),
	// 				ac.FormatMoney(trade.Trade.TotalRub),
	// 				trade.Trade.Percent,
	// 				trade.Trade.PercentRub))

	// 		}

	// 		sb.WriteString("\n")

	// 	}

	// 	sb.WriteString(fmt.Sprintf(texts.AccountTradesTotalDay, ac.FormatMoney(totalDayRub)))

	// 	message = tgbotapi.NewMessage(user.ChatID, "")
	// 	message.Text = sb.String()

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	// message = tgbotapi.NewMessage(user.ChatID, "")
	// message.Text = fmt.Sprintf(texts.AccountTradesTotal, title, ac.FormatMoney(totalRub))

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil
}
