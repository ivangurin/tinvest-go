package bot

import (
	"context"
	"fmt"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountTrades(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccountTrades.FindAllStringSubmatch(request.Text, -1)
	if len(result) == 0 || len(result[0]) == 0 {
		return fmt.Errorf("can't parse account id in '%s'", request.Text)
	}
	accountID := result[0][1]

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

	message := tgbotapi.NewMessage(request.Chat.ID, "")
	message.ReplyMarkup =
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesCurrDay), fmt.Sprintf(commandAccountTradesFor, account.ID, tradesForCurrDay)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesPrevDay), fmt.Sprintf(commandAccountTradesFor, account.ID, tradesForPrevDay)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesCurrWeek), fmt.Sprintf(commandAccountTradesFor, account.ID, tradesForCurrWeek)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesPrevWeek), fmt.Sprintf(commandAccountTradesFor, account.ID, tradesForPrevWeek)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesCurrMonth), fmt.Sprintf(commandAccountTradesFor, account.ID, tradesForCurrMonth)),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf(texts.AccountTradesPrevMonth), fmt.Sprintf(commandAccountTradesFor, account.ID, tradesForPrevMonth)),
			))

	_, err = a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
	}

	return nil
}
