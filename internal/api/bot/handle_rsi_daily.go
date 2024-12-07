package bot

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"
	"tinvest-go/internal/indicators"
	"tinvest-go/internal/model"
	contractv1 "tinvest-go/internal/pb"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleRSIDaily(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
	messageID, err := a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.Processing)
	if err != nil {
		return err
	}
	defer func() {
		err := a.botClient.DeleteMessage(ctx, request.Chat.ID, messageID)
		if err != nil {
			logger.Errorf(ctx, "error on delete message: %s", err.Error())
		}
	}()

	type row struct {
		InstrumentID string
		Ticker       string
		PrevRsiValue float64
		LastRsiValue float64
		LastPrice    float64
		Currency     string
	}

	instruments, err := a.tinvestService.GetFavorites(ctx, user.Token)
	if err != nil {
		return fmt.Errorf("error on get favorite instruments: %w", err)
	}

	from := time.Now().AddDate(0, 0, -365)
	to := time.Now()

	result := make([]row, 0, len(instruments))

	for _, instrument := range instruments {
		if !instrument.Trading {
			continue
		}

		candles, err := a.tinvestService.GetCandles(ctx, user.Token, instrument.ID, contractv1.CandleInterval_CANDLE_INTERVAL_DAY, from, to)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("error on get candles for %+v: %s", *instrument, err.Error()))
			continue
		}

		rsi := indicators.GetRSI(candles)

		prevRsiValue := rsi[len(rsi)-2].Value
		lastRsiValue := rsi[len(rsi)-1].Value

		if lastRsiValue > 35 && lastRsiValue < 65 {
			continue
		}

		row :=
			row{
				Ticker:       instrument.Ticker,
				PrevRsiValue: prevRsiValue,
				LastRsiValue: lastRsiValue,
				LastPrice:    instrument.LastPrice,
				Currency:     instrument.Currency,
			}

		result = append(result, row)

	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].LastRsiValue > result[j].LastRsiValue
	})

	mb := strings.Builder{}

	mb.WriteString(texts.RSIDTitle)
	mb.WriteString("\n")
	mb.WriteString(texts.RSIDDescription)
	mb.WriteString("\n")

	for _, resultRow := range result {
		if resultRow.LastRsiValue >= 70 {
			mb.WriteString(
				fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
					resultRow.Ticker,
					resultRow.LastRsiValue,
					resultRow.PrevRsiValue,
					a.accounting.FormatMoney(resultRow.LastPrice),
					resultRow.Currency))
		}
	}

	mb.WriteString("...")
	mb.WriteString("\n")

	for _, resultRow := range result {
		if resultRow.LastRsiValue >= 65 && resultRow.LastRsiValue < 70 {
			mb.WriteString(
				fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
					resultRow.Ticker,
					resultRow.LastRsiValue,
					resultRow.PrevRsiValue,
					a.accounting.FormatMoney(resultRow.LastPrice),
					resultRow.Currency))
		}
	}

	mb.WriteString("...")
	mb.WriteString("\n")

	for _, resultRow := range result {
		if resultRow.LastRsiValue > 30 && resultRow.LastRsiValue < 35 {
			mb.WriteString(
				fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
					resultRow.Ticker,
					resultRow.LastRsiValue,
					resultRow.PrevRsiValue,
					a.accounting.FormatMoney(resultRow.LastPrice),
					resultRow.Currency))
		}
	}

	mb.WriteString("...")
	mb.WriteString("\n")

	for _, resultRow := range result {
		if resultRow.LastRsiValue <= 30 {
			mb.WriteString(
				fmt.Sprintf("<b>%s</b>: %.2f(%.2f), Цена: %s %s\n",
					resultRow.Ticker,
					resultRow.LastRsiValue,
					resultRow.PrevRsiValue,
					a.accounting.FormatMoney(resultRow.LastPrice),
					resultRow.Currency))
		}
	}

	_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, mb.String())
	if err != nil {
		return err
	}

	return nil
}
