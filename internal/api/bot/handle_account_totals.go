package bot

import (
	"context"
	"fmt"
	"strings"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	tinvest_service "tinvest-go/internal/service/tinvest"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountTotals(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccountTotals.FindAllStringSubmatch(request.Text, -1)
	if len(result) == 0 || len(result[0]) == 0 {
		return fmt.Errorf("can't parse account id in '%s'", request.Text)
	}
	accountID := result[0][1]

	account, err := a.tinvestService.GetAccountByID(ctx, user.Token, accountID)
	if err != nil {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.GetDataError)
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

	totals, err := a.tinvestService.GetTotals(ctx, user.Token, accountID, tinvest_service.From, time.Now())
	if err != nil {
		return err
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(texts.AccountTotalsTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
	sb.WriteString("\n")

	totalRub := float64(0)
	for _, total := range totals {
		totalRub += total.TotalRub
		sb.WriteString(fmt.Sprintf("<b>%s</b>: %s(%s руб.)\n", strings.ToUpper(total.Currency), a.accounting.FormatMoney(total.Total), a.accounting.FormatMoney(total.TotalRub)))
	}

	sb.WriteString(fmt.Sprintf(texts.TotalRubles, a.accounting.FormatMoney(totalRub)))

	_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, sb.String())
	if err != nil {
		return err
	}

	return nil
}
