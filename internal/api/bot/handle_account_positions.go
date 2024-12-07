package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const positionsPerMessage = 50

func (a *api) HandleAccountPositions(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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
	// 	return errors.Errorf("error on get tiClient: %+v", err)
	// }

	// account, err := tiClient.GetAccount(accountId)
	// if err != nil {
	// 	return errors.Errorf("error on get account %s: %+v", accountId, err)
	// }

	// if account == nil {
	// 	return nil
	// }

	// tiClient.SetAccountId(account.Id)

	// positions, err := tiClient.GetPositions(time.Now(), "")
	// if err != nil {
	// 	return errors.Errorf("error on get positions for %s: %+v", accountId, err)
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
	// 		sb.WriteString(fmt.Sprintf(texts.AccountPositionsTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
	// 		sb.WriteString("\n")
	// 	}

	// 	if messageCounter > 1 {
	// 		sb.WriteString(fmt.Sprintf(texts.AccountPositionsPart, i+1, int(messageCounter)))
	// 		sb.WriteString("\n")
	// 	}

	// 	for _, position := range positions[fromPosition:toPosition] {

	// 		if position.Ticker == "" {
	// 			continue
	// 		}

	// 		sb.WriteString(
	// 			fmt.Sprintf(texts.PositionResult,
	// 				position.Ticker,
	// 				a.accounting.FormatMoney(position.Total),
	// 				position.Currency,
	// 				a.accounting.FormatMoney(position.TotalRub)))

	// 	}

	// 	message = tgbotapi.NewMessage(user.ChatID, sb.String())

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	return nil
}
