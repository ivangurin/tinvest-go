package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountTrades(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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

	// message = tgbotapi.NewMessage(user.ChatID, "")
	// message.Text = texts.AccountTradesFor

	// message.ReplyMarkup =
	// 	tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesCurrDay), fmt.Sprintf("/%s/%s/%s/currDay", collectionAccounts, accountId, actionTrades)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesPrevDay), fmt.Sprintf("/%s/%s/%s/prevDay", collectionAccounts, accountId, actionTrades)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesCurrWeek), fmt.Sprintf("/%s/%s/%s/currWeek", collectionAccounts, accountId, actionTrades)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesPrevWeek), fmt.Sprintf("/%s/%s/%s/prevWeek", collectionAccounts, accountId, actionTrades)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesCurrMonth), fmt.Sprintf("/%s/%s/%s/currMonth", collectionAccounts, accountId, actionTrades)),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesPrevMonth), fmt.Sprintf("/%s/%s/%s/prevMonth", collectionAccounts, accountId, actionTrades)),
	// 		))

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil
}
