package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountTotals(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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
	// 	return nil
	// }

	// tiClient.SetAccountId(account.Id)

	// totals, err := tiClient.GetTotals(time.Now())
	// if err != nil {
	// 	return err
	// }

	// sb := strings.Builder{}

	// sb.WriteString(fmt.Sprintf(texts.AccountTotalsTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
	// sb.WriteString("\n")

	// totalRub := float64(0)

	// for _, total := range totals {
	// 	totalRub += total.TotalRub
	// 	sb.WriteString(fmt.Sprintf("<b>%s</b>: %s(%s руб.)\n", strings.ToUpper(total.Currency), a.accounting.FormatMoney(total.Total), a.accounting.FormatMoney(total.TotalRub)))
	// }

	// sb.WriteString(fmt.Sprintf(texts.TotalRubles, a.accounting.FormatMoney(totalRub)))

	// message = tgbotapi.NewMessage(user.ChatID, "")
	// message.ParseMode = tgbotapi.ModeHTML
	// message.Text = sb.String()

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil
}
