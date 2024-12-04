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

	// 	sb.WriteString(fmt.Sprintf(texts.PositionPriceBuy, ac.FormatMoney(position.PriceBuy), ac.FormatMoney(position.PriceBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionValueBuy, ac.FormatMoney(position.ValueBuy), ac.FormatMoney(position.ValueBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionNKDBuy, ac.FormatMoney(position.NKDBuy), ac.FormatMoney(position.NKDBuyRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionQuantitySell, position.QuantitySell))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionPriceSell, ac.FormatMoney(position.PriceSell), ac.FormatMoney(position.PriceSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionValueSell, ac.FormatMoney(position.ValueSell), ac.FormatMoney(position.ValueSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionNKDSell, ac.FormatMoney(position.NKDSell), ac.FormatMoney(position.NKDSellRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionQuantityEnd, position.QuantityEnd))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionPriceEnd, ac.FormatMoney(position.PriceEnd), ac.FormatMoney(position.PriceEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionValueEnd, ac.FormatMoney(position.ValueEnd), ac.FormatMoney(position.ValueEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionNKDEnd, ac.FormatMoney(position.NKDEnd), ac.FormatMoney(position.NKDEndRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionDividends, ac.FormatMoney(position.Dividents), ac.FormatMoney(position.DividentsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionCoupons, ac.FormatMoney(position.Coupons), ac.FormatMoney(position.CouponsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionOvernight, ac.FormatMoney(position.Overnight), ac.FormatMoney(position.OvernightRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionTaxes, ac.FormatMoney(position.Taxes), ac.FormatMoney(position.TaxesRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionCommissions, ac.FormatMoney(position.Commissions), ac.FormatMoney(position.CommissionsRub)))
	// 	sb.WriteString("\n")

	// 	sb.WriteString("\n")

	// 	sb.WriteString(fmt.Sprintf(texts.PositionTotal, ac.FormatMoney(position.Total), ac.FormatMoney(position.TotalRub)))
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
