package bot

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// nolint:funlen
func (a *api) HandleAccountTradesFor(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccountTradesFor.FindAllStringSubmatch(request.Text, -1)
	if len(result) == 0 || len(result[0]) == 0 {
		return fmt.Errorf("can't parse account id in '%s'", request.Text)
	}
	accountID := result[0][1]
	tradesFor := result[0][2]

	account, err := a.tinvestService.GetAccountByID(ctx, user.Token, accountID)
	if err != nil {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.AccountNotFound, accountID))
		if err != nil {
			return err
		}
		return err
	}

	if account == nil {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.AccountNotFound, accountID))
		if err != nil {
			return err
		}
		return nil
	}

	now := time.Now()
	weekday := now.Weekday()
	if weekday == 0 {
		weekday = 7
	}

	var title string
	var from time.Time
	var to time.Time
	switch tradesFor {
	case tradesForCurrDay:
		title = texts.AccountTradesCurrDay
		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	case tradesForPrevDay:
		title = texts.AccountTradesPrevDay
		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		from = from.AddDate(0, 0, -1)
		to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
		to = to.AddDate(0, 0, -1)
	case tradesForCurrWeek:
		title = texts.AccountTradesCurrWeek
		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		from = from.AddDate(0, 0, -int(weekday-1))
		to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	case tradesForPrevWeek:
		title = texts.AccountTradesPrevWeek
		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		from = from.AddDate(0, 0, -int(7+weekday-1))
		to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
		to = to.AddDate(0, 0, -int(weekday))
	case tradesForCurrMonth:
		title = texts.AccountTradesCurrMonth
		from = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		to = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	case tradesForPrevMonth:
		title = texts.AccountTradesPrevMonth
		to = time.Date(now.Year(), now.Month(), 1, 23, 59, 59, 0, time.UTC)
		to = to.AddDate(0, 0, -1)
		from = time.Date(to.Year(), to.Month(), 1, 0, 0, 0, 0, time.UTC)
	default:
		return nil
	}

	trades, err := a.tinvestService.GetTrades(ctx, user.Token, account.ID, from, to)
	if err != nil {
		return err
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(texts.AccountTradesTitle, title, from.Format("02.01.2006"), to.Format("02.01.2006"), account.Name))
	sb.WriteString("\n")

	if len(trades) == 0 {
		sb.WriteString(texts.AccountTradesNoTrades)
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, sb.String())
		if err != nil {
			return err
		}
		return nil
	}

	group := map[time.Time]model.Trades{}
	for _, trade := range trades {
		date := time.Date(trade.Time.Year(), trade.Time.Month(), trade.Time.Day(), 0, 0, 0, 0, time.UTC)
		group[date] = append(group[date], trade)
	}

	keys := []time.Time{}
	for key := range group {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	totalRub := float64(0)
	for keyIndex, key := range keys {
		if keyIndex > 0 {
			sb = strings.Builder{}
		}

		sb.WriteString(fmt.Sprintf("<b>%s:</b>\n", key.Format("02.01.2006")))

		trades = group[key]
		totalDayRub := float64(0)

		for _, trade := range trades {
			totalRub += trade.Trade.TotalRub
			totalDayRub += trade.Trade.TotalRub
			if trade.Currency == model.CurrencyRUB {
				sb.WriteString(fmt.Sprintf(texts.AccountTradesItemRub,
					trade.Ticker,
					trade.Trade.QuantityBuy,
					trade.Currency,
					a.accounting.FormatMoney(trade.Trade.ValueBuy),
					a.accounting.FormatMoney(trade.Trade.ValueSell),
					a.accounting.FormatMoney(trade.Trade.Total),
					trade.Trade.Percent))
			} else {
				sb.WriteString(fmt.Sprintf(texts.AccountTradesItem,
					trade.Ticker,
					trade.Trade.QuantityBuy,
					trade.Currency,
					a.accounting.FormatMoney(trade.Trade.ValueBuy),
					a.accounting.FormatMoney(trade.Trade.ValueBuyRub),
					a.accounting.FormatMoney(trade.Trade.ValueSell),
					a.accounting.FormatMoney(trade.Trade.ValueSellRub),
					a.accounting.FormatMoney(trade.Trade.Total),
					a.accounting.FormatMoney(trade.Trade.TotalRub),
					trade.Trade.Percent,
					trade.Trade.PercentRub))
			}
			sb.WriteString("\n")
		}

		sb.WriteString(fmt.Sprintf(texts.AccountTradesTotalDay, a.accounting.FormatMoney(totalDayRub)))

		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, sb.String())
		if err != nil {
			return err
		}
	}

	_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.AccountTradesTotal, title, a.accounting.FormatMoney(totalRub)))
	if err != nil {
		return err
	}

	return nil
}
