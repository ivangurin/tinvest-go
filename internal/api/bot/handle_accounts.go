package bot

import (
	"context"
	"fmt"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccounts(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	accounts, err := a.tinvestService.GetAccounts(ctx, user.Token)
	if err != nil {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.GetDataError)
		if err != nil {
			return err
		}
		return err
	}

	rows := make([][]tgbotapi.InlineKeyboardButton, 0, len(accounts))
	for _, account := range accounts {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				account.Name,
				fmt.Sprintf(commandAccount, account.ID)))
		rows = append(rows, row)
	}

	message := tgbotapi.NewMessage(request.Chat.ID, texts.AccountsList)
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	_, err = a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
	}

	return nil
}
