package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountPortfolio(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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
	// 	return errors.Errorf("error on create tinkoff client: %+v", err)
	// }

	// account, err := tiClient.GetAccount(accountId)
	// if err != nil {
	// 	return errors.Errorf("error on get account %s: %+v", accountId, err)
	// }

	// if account == nil {
	// 	return nil
	// }

	// tiClient.SetAccountId(account.Id)

	// positions, err := tiClient.GetPortfolio()
	// if err != nil {
	// 	return errors.Errorf("error on get portfolio for %s: %+v", accountId, err)
	// }

	// messageCounter, rest := math.Modf(float64(len(positions)) / float64(positionsPerMessage))

	// if rest > 0 {
	// 	messageCounter += 1
	// }

	// var sb strings.Builder
	// var fromPosition = 0
	// var toPosition = 0

	// for i := 0; i < int(messageCounter); i++ {

	// 	fromPosition = i * positionsPerMessage
	// 	toPosition = fromPosition + positionsPerMessage

	// 	if toPosition > len(positions) {
	// 		toPosition = len(positions)
	// 	}

	// 	sb = strings.Builder{}

	// 	if i == 0 {
	// 		sb.WriteString(fmt.Sprintf(texts.AccountPortfolioTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
	// 		sb.WriteString("\n")
	// 	}

	// 	if messageCounter > 1 {
	// 		sb.WriteString(fmt.Sprintf(texts.AccountPositionsPart, i+1, int(messageCounter)))
	// 		sb.WriteString("\n")
	// 	}

	// 	for _, position := range positions[fromPosition:toPosition] {

	// 		if position.Currency == tinvest.CurrencyRUB {

	// 			sb.WriteString(
	// 				fmt.Sprintf(texts.AccountPortfolioPositionRub,
	// 					position.Ticker,
	// 					position.Quantity,
	// 					position.Currency,
	// 					ac.FormatMoney(position.Price),
	// 					ac.FormatMoney(position.Value),
	// 					ac.FormatMoney(position.PriceEnd),
	// 					ac.FormatMoney(position.ValueEnd),
	// 					ac.FormatMoney(position.Total),
	// 					ac.FormatMoney(position.Percent)))

	// 		} else {

	// 			sb.WriteString(
	// 				fmt.Sprintf(texts.AccountPortfolioPosition,
	// 					position.Ticker,
	// 					position.Quantity,
	// 					position.Currency,
	// 					ac.FormatMoney(position.Price),
	// 					ac.FormatMoney(position.Value),
	// 					ac.FormatMoney(position.ValueRub),
	// 					ac.FormatMoney(position.PriceEnd),
	// 					ac.FormatMoney(position.ValueEnd),
	// 					ac.FormatMoney(position.ValueEndRub),
	// 					ac.FormatMoney(position.Total),
	// 					ac.FormatMoney(position.Percent),
	// 					ac.FormatMoney(position.TotalRub),
	// 					ac.FormatMoney(position.PercentRub)))

	// 		}

	// 	}

	// 	message = tgbotapi.NewMessage(user.ChatID, "")
	// 	message.ParseMode = tgbotapi.ModeHTML
	// 	message.Text = sb.String()

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	return nil
}
