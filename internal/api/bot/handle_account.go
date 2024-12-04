package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccount(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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
	// 	return err
	// }

	// account, err := tiClient.GetAccount(accountId)
	// if err != nil {
	// 	return err
	// }

	// if account == nil {

	// 	message = tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(texts.AccountNotFound, accountId))

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// }

	// message = tgbotapi.NewMessage(user.ChatID, "")
	// message.Text = fmt.Sprintf(texts.AccountTitle, account.Name)
	// message.ReplyMarkup =
	// 	tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(texts.AccountPortfolio, fmt.Sprintf("/%s/%s/%s", collectionAccounts, accountId, actionPortfolio)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(texts.AccountTotals, fmt.Sprintf("/%s/%s/%s", collectionAccounts, accountId, actionTotals)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(texts.AccountDetail, fmt.Sprintf("/%s/%s/%s", collectionAccounts, accountId, actionDetail)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(texts.AccountPositions, fmt.Sprintf("/%s/%s/%s", collectionAccounts, accountId, actionPositions)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(texts.AccountPosition, fmt.Sprintf("/%s/%s/%s", collectionAccounts, accountId, actionPosition)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(texts.AccountTrades, fmt.Sprintf("/%s/%s/%s", collectionAccounts, accountId, actionTrades)),
	// 		))

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil

}
