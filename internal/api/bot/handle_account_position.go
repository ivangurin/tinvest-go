package bot

import (
	"context"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountPosition(ctx context.Context, user *model.User, request *tgbotapi.Message) error {

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

	// 	message = tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(texts.AccountNotFound, accountId))

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// }

	// if ticker == "" {

	// 	message = tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(texts.AccountPositionHelp, account.Name))

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil

	// }

	// instrument, err := tiClient.GetInstrumentByTicker(strings.ToUpper(ticker))
	// if err != nil {
	// 	return errors.Errorf("error on get instrument by ticker for %s: %+v", ticker, err)
	// }

	// if instrument == nil {

	// 	message = tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(texts.InstrumentNotFound, ticker))

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil

	// }

	// tiClient.SetAccountId(account.Id)

	// positions, err := tiClient.GetPositions(time.Now(), instrument.Figi)
	// if err != nil {
	// 	return errors.Errorf("error on get position for %s/%s: %+v", accountId, instrument.Figi, err)
	// }

	// if len(positions) == 0 {

	// 	message = tgbotapi.NewMessage(user.ChatID, fmt.Sprintf(texts.PositionNotFound, instrument.Ticker))

	// 	_, err = a.botClient.SendMessage(ctx, &message)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil

	// }

	// sb := strings.Builder{}

	// for _, position := range positions {

	// 	sb.WriteString(fmt.Sprintf(texts.AccountPositionDetailTitle, position.Ticker, account.Name, time.Now().Format("02.01.2006 15:04")))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionCurrency, strings.ToUpper(position.Currency)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionQuantityBuy, position.QuantityBuy))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionPriceBuy, a.accounting.FormatMoney(position.PriceBuy), a.accounting.FormatMoney(position.PriceBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionValueBuy, a.accounting.FormatMoney(position.ValueBuy), a.accounting.FormatMoney(position.ValueBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionNKDBuy, a.accounting.FormatMoney(position.NKDBuy), a.accounting.FormatMoney(position.NKDBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionQuantitySell, position.QuantitySell))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionPriceSell, a.accounting.FormatMoney(position.PriceSell), a.accounting.FormatMoney(position.PriceSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionValueSell, a.accounting.FormatMoney(position.ValueSell), a.accounting.FormatMoney(position.ValueSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionNKDSell, a.accounting.FormatMoney(position.NKDSell), a.accounting.FormatMoney(position.NKDSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionQuantityEnd, position.QuantityEnd))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionPriceEnd, a.accounting.FormatMoney(position.PriceEnd), a.accounting.FormatMoney(position.PriceEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionValueEnd, a.accounting.FormatMoney(position.ValueEnd), a.accounting.FormatMoney(position.ValueEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionNKDEnd, a.accounting.FormatMoney(position.NKDEnd), a.accounting.FormatMoney(position.NKDEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionDividends, a.accounting.FormatMoney(position.Dividents), a.accounting.FormatMoney(position.DividentsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionCoupons, a.accounting.FormatMoney(position.Coupons), a.accounting.FormatMoney(position.CouponsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionOvernight, a.accounting.FormatMoney(position.Overnight), a.accounting.FormatMoney(position.OvernightRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionTaxes, a.accounting.FormatMoney(position.Taxes), a.accounting.FormatMoney(position.TaxesRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionCommissions, a.accounting.FormatMoney(position.Commissions), a.accounting.FormatMoney(position.CommissionsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionTotal, a.accounting.FormatMoney(position.Total), a.accounting.FormatMoney(position.TotalRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// }

	// message = tgbotapi.NewMessage(user.ChatID, sb.String())

	// _, err = a.botClient.SendMessage(ctx, &message)
	// if err != nil {
	// 	return err
	// }

	return nil
}
