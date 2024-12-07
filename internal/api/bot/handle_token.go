package bot

import (
	"context"
	"fmt"
	"tinvest-go/internal/model"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleToken(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
	if isCommand(request.Text) {
		_, err := a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.WriteToken)
		if err != nil {
			return err
		}
	} else {
		err := a.botClient.DeleteMessage(ctx, request.Chat.ID, request.MessageID)
		if err != nil {
			return err
		}

		var create bool
		if user == nil {
			create = true
			user = &model.User{
				ID:     request.From.ID,
				ChatID: request.Chat.ID,
			}
		}

		user.Username = request.From.UserName
		user.FirstName = request.From.FirstName
		user.LastName = request.From.LastName
		user.Token = request.Text

		if create {
			user, err = a.userService.CreateUser(ctx, user)
			if err != nil {
				return fmt.Errorf("failed to create user %d: %w", user.ID, err)
			}
		} else {
			err = a.userService.UpdateUser(ctx, user)
			if err != nil {
				return fmt.Errorf("failed to update user %d: %w", user.ID, err)
			}
		}

		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.TokenIsDeleted)
		if err != nil {
			return err
		}

		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.TokenIsSaved)
		if err != nil {
			return err
		}

		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, texts.DoCommands)
		if err != nil {
			return err
		}
	}

	return nil
}
