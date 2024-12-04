package bot

import (
	"context"
	"fmt"
	"tinvest-go/internal/model"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleStart(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
	message := tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(texts.Hi, user.FirstName))
	_, err := a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
	}

	message = tgbotapi.NewMessage(user.ChatID, texts.About)
	_, err = a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
	}

	message = tgbotapi.NewMessage(user.ChatID, texts.YouShouldDoThis)
	message.ReplyMarkup =
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL(texts.GetTkn, texts.TokenUrl)),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.SetTkn, commandToken)),
		)

	_, err = a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
	}

	return nil
}
