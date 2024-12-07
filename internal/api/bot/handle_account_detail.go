package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountDetail(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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

	// sb.WriteString(fmt.Sprintf(texts.AccountDetailTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
	// sb.WriteString("\n")

	// totalRub := float64(0)

	// for _, total := range totals {

	// 	totalRub += total.TotalRub

	// 	sb.WriteString(fmt.Sprintf(texts.CurrecnyLabel, strings.ToUpper(total.Currency)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.ValueBuyLabel, a.accounting.FormatMoney(total.ValueBuy), a.accounting.FormatMoney(total.ValueBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.ValueSellLabel, a.accounting.FormatMoney(total.ValueSell), a.accounting.FormatMoney(total.ValueSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.ValueEndLabel, a.accounting.FormatMoney(total.ValueEnd), a.accounting.FormatMoney(total.ValueEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.NKDBuyLabel, a.accounting.FormatMoney(total.NKDBuy), a.accounting.FormatMoney(total.NKDBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.NKDSellLabel, a.accounting.FormatMoney(total.NKDSell), a.accounting.FormatMoney(total.NKDSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.NKDEndLabel, a.accounting.FormatMoney(total.NKDEnd), a.accounting.FormatMoney(total.NKDEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CouponsLabel, a.accounting.FormatMoney(total.Coupons), a.accounting.FormatMoney(total.CouponsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.DividentsLabel, a.accounting.FormatMoney(total.Dividents), a.accounting.FormatMoney(total.DividentsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.OvernightLabel, a.accounting.FormatMoney(total.Overnight), a.accounting.FormatMoney(total.OvernightRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.TaxesLabel, a.accounting.FormatMoney(total.Taxes), a.accounting.FormatMoney(total.TaxesRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CommissionsLabel, a.accounting.FormatMoney(total.Commissions), a.accounting.FormatMoney(total.CommissionsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CommissionsTrackLabel, a.accounting.FormatMoney(total.TrackFee), a.accounting.FormatMoney(total.TrackFeeRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CommissionsResultLabel, a.accounting.FormatMoney(total.ResultFee), a.accounting.FormatMoney(total.ResultFeeRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.TotalLabel, a.accounting.FormatMoney(total.Total), a.accounting.FormatMoney(total.TotalRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// }

	// sb.WriteString(fmt.Sprintf(texts.TotalRubLabel, a.accounting.FormatMoney(totalRub)))
	// sb.WriteString("\n")

	// message = tgbotapi.NewMessage(user.ChatID, "")
	// message.ParseMode = tgbotapi.ModeHTML
	// message.Text = sb.String()

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil
}
