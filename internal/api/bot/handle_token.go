package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleToken(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

	// message := tgbotapi.NewMessage(user.ChatID, texts.WriteToken)

	// _, err = a.botClient.SendMessage(message)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// func handleTokenResponse(user *model.User, message *tgbotapi.Message) error {

// 	err = deleteMessage(user.ChatID, message.MessageID)
// 	if err != nil {
// 		return err
// 	}

// 	user.Token = message.Text

// 	err = repo.UpdateUser(user)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = sendMessageWithText(message.Chat.ID, texts.TokenIsDeleted)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = sendMessageWithText(message.Chat.ID, texts.TokenIsSaved)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = sendMessageWithText(message.Chat.ID, texts.DoCommands)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
