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

	// 	sb.WriteString(fmt.Sprintf(texts.ValueBuyLabel, ac.FormatMoney(total.ValueBuy), ac.FormatMoney(total.ValueBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.ValueSellLabel, ac.FormatMoney(total.ValueSell), ac.FormatMoney(total.ValueSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.ValueEndLabel, ac.FormatMoney(total.ValueEnd), ac.FormatMoney(total.ValueEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.NKDBuyLabel, ac.FormatMoney(total.NKDBuy), ac.FormatMoney(total.NKDBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.NKDSellLabel, ac.FormatMoney(total.NKDSell), ac.FormatMoney(total.NKDSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.NKDEndLabel, ac.FormatMoney(total.NKDEnd), ac.FormatMoney(total.NKDEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CouponsLabel, ac.FormatMoney(total.Coupons), ac.FormatMoney(total.CouponsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.DividentsLabel, ac.FormatMoney(total.Dividents), ac.FormatMoney(total.DividentsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.OvernightLabel, ac.FormatMoney(total.Overnight), ac.FormatMoney(total.OvernightRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.TaxesLabel, ac.FormatMoney(total.Taxes), ac.FormatMoney(total.TaxesRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CommissionsLabel, ac.FormatMoney(total.Commissions), ac.FormatMoney(total.CommissionsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CommissionsTrackLabel, ac.FormatMoney(total.TrackFee), ac.FormatMoney(total.TrackFeeRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.CommissionsResultLabel, ac.FormatMoney(total.ResultFee), ac.FormatMoney(total.ResultFeeRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.TotalLabel, ac.FormatMoney(total.Total), ac.FormatMoney(total.TotalRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// }

	// sb.WriteString(fmt.Sprintf(texts.TotalRubLabel, ac.FormatMoney(totalRub)))
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
