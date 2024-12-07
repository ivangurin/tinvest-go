package bot

import (
	"context"
	"fmt"
	"tinvest-go/internal/model"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleStart(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
	_, err := a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.Hi, request.From.FirstName))
	if err != nil {
		return err
	}

	_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.About)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(request.Chat.ID, texts.YouShouldDoThis)
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
