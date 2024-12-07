package bot

import (
	"context"
	"fmt"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccount(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccount.FindAllStringSubmatch(request.Text, -1)
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

	message := tgbotapi.NewMessage(request.Chat.ID, fmt.Sprintf(texts.AccountTitle, account.Name))
	message.ReplyMarkup =
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.AccountPortfolio, fmt.Sprintf(commandAccountPortfolio, account.ID)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.AccountTotals, fmt.Sprintf(commandAccountTotals, account.ID)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.AccountDetail, fmt.Sprintf(commandAccountDetail, account.ID)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.AccountPositions, fmt.Sprintf(commandAccountPositions, account.ID)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.AccountPosition, fmt.Sprintf(commandAccountPosition, account.ID)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(texts.AccountTrades, fmt.Sprintf(commandAccountTrades, account.ID)),
			))

	_, err = a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
	}

	return nil
}
