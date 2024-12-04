package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleRSIDaily(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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

	// tiClient, err := tinvest.NewClient(user.Token, repo)
	// if err != nil {
	// 	return errors.Errorf("error on get tiClient: %+v", err)
	// }

	// type resultRow struct {
	// 	Ticker       string
	// 	PrevRsiValue float64
	// 	LastRsiValue float64
	// 	LastPrice    float64
	// 	Currency     string
	// }

	// instruments, err := tiClient.GetFavorites()
	// if err != nil {
	// 	return errors.Errorf("error on get favorite instruments: %+v", err)
	// }

	// from := time.Now().AddDate(0, 0, -365)
	// to := time.Now()

	// result := make([]resultRow, 0, len(instruments))

	// for figi, instrument := range instruments {

	// 	if !instrument.Trading {
	// 		continue
	// 	}

	// 	candles, err := tiClient.GetCandles(figi, contractv1.CandleInterval_CANDLE_INTERVAL_DAY, from, to)
	// 	if err != nil {
	// 		logger.Error(errors.Errorf("error on get candles for %s: %+v\n", figi, err))
	// 		continue
	// 	}

	// 	rsi := indicators.GetRSI(candles)

	// 	prevRsiValue := rsi[len(rsi)-2].Value
	// 	lastRsiValue := rsi[len(rsi)-1].Value

	// 	if lastRsiValue > 35 && lastRsiValue < 65 {
	// 		continue
	// 	}

	// 	resultRow :=
	// 		resultRow{
	// 			Ticker:       instrument.Ticker,
	// 			PrevRsiValue: prevRsiValue,
	// 			LastRsiValue: lastRsiValue,
	// 			LastPrice:    instrument.LastPrice,
	// 			Currency:     instrument.Currency,
	// 		}

	// 	result = append(result, resultRow)

	// }

	// sort.Slice(result, func(i, j int) bool {
	// 	return result[i].LastRsiValue > result[j].LastRsiValue
	// })

	// mb := strings.Builder{}

	// mb.WriteString(texts.RSIDTitle)
	// mb.WriteString("\n")
	// mb.WriteString(texts.RSIDDescription)
	// mb.WriteString("\n")

	// for _, resultRow := range result {
	// 	if resultRow.LastRsiValue >= 70 {
	// 		mb.WriteString(
	// 			fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
	// 				resultRow.Ticker,
	// 				resultRow.LastRsiValue,
	// 				resultRow.PrevRsiValue,
	// 				ac.FormatMoney(resultRow.LastPrice),
	// 				resultRow.Currency))
	// 	}
	// }

	// mb.WriteString("...")
	// mb.WriteString("\n")

	// for _, resultRow := range result {
	// 	if resultRow.LastRsiValue >= 65 && resultRow.LastRsiValue < 70 {
	// 		mb.WriteString(
	// 			fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
	// 				resultRow.Ticker,
	// 				resultRow.LastRsiValue,
	// 				resultRow.PrevRsiValue,
	// 				ac.FormatMoney(resultRow.LastPrice),
	// 				resultRow.Currency))
	// 	}
	// }

	// mb.WriteString("...")
	// mb.WriteString("\n")

	// for _, resultRow := range result {
	// 	if resultRow.LastRsiValue > 30 && resultRow.LastRsiValue < 35 {
	// 		mb.WriteString(
	// 			fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
	// 				resultRow.Ticker,
	// 				resultRow.LastRsiValue,
	// 				resultRow.PrevRsiValue,
	// 				ac.FormatMoney(resultRow.LastPrice),
	// 				resultRow.Currency))
	// 	}
	// }

	// mb.WriteString("...")
	// mb.WriteString("\n")

	// for _, resultRow := range result {
	// 	if resultRow.LastRsiValue <= 30 {
	// 		mb.WriteString(
	// 			fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
	// 				resultRow.Ticker,
	// 				resultRow.LastRsiValue,
	// 				resultRow.PrevRsiValue,
	// 				ac.FormatMoney(resultRow.LastPrice),
	// 				resultRow.Currency))
	// 	}
	// }

	// message = tgbotapi.NewMessage(user.ChatID, "")

	// message.ParseMode = "HTML"
	// message.Text = mb.String()

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil
}
