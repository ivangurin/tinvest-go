package bot

import (
	"context"
	"fmt"
	"strings"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	tinvest_service "tinvest-go/internal/service/tinvest"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountDetail(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccountDetail.FindAllStringSubmatch(request.Text, -1)
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

	totals, err := a.tinvestService.GetTotals(ctx, user.Token, accountID, tinvest_service.From, time.Now())
	if err != nil {
		return err
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(texts.AccountDetailTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
	sb.WriteString("\n")

	totalRub := float64(0)
	for _, total := range totals {
		totalRub += total.TotalRub

		sb.WriteString(fmt.Sprintf(texts.CurrencyLabel, strings.ToUpper(total.Currency)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.ValueBuyLabel, a.accounting.FormatMoney(total.ValueBuy), a.accounting.FormatMoney(total.ValueBuyRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.ValueSellLabel, a.accounting.FormatMoney(total.ValueSell), a.accounting.FormatMoney(total.ValueSellRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.ValueEndLabel, a.accounting.FormatMoney(total.ValueEnd), a.accounting.FormatMoney(total.ValueEndRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.NKDBuyLabel, a.accounting.FormatMoney(total.NKDBuy), a.accounting.FormatMoney(total.NKDBuyRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.NKDSellLabel, a.accounting.FormatMoney(total.NKDSell), a.accounting.FormatMoney(total.NKDSellRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.NKDEndLabel, a.accounting.FormatMoney(total.NKDEnd), a.accounting.FormatMoney(total.NKDEndRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.CouponsLabel, a.accounting.FormatMoney(total.Coupons), a.accounting.FormatMoney(total.CouponsRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.DividendsLabel, a.accounting.FormatMoney(total.Dividends), a.accounting.FormatMoney(total.DividendsRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.OvernightLabel, a.accounting.FormatMoney(total.Overnight), a.accounting.FormatMoney(total.OvernightRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.TaxesLabel, a.accounting.FormatMoney(total.Taxes), a.accounting.FormatMoney(total.TaxesRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.CommissionsLabel, a.accounting.FormatMoney(total.Commissions), a.accounting.FormatMoney(total.CommissionsRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.CommissionsTrackLabel, a.accounting.FormatMoney(total.TrackFee), a.accounting.FormatMoney(total.TrackFeeRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.CommissionsResultLabel, a.accounting.FormatMoney(total.ResultFee), a.accounting.FormatMoney(total.ResultFeeRub)))
		sb.WriteString("\n")

		sb.WriteString(fmt.Sprintf(texts.TotalLabel, a.accounting.FormatMoney(total.Total), a.accounting.FormatMoney(total.TotalRub)))
		sb.WriteString("\n")

		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf(texts.TotalRubLabel, a.accounting.FormatMoney(totalRub)))
	sb.WriteString("\n")

	_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, sb.String())
	if err != nil {
		return err
	}

	return nil
}
