package bot

import (
	"context"
	"fmt"
	"strings"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	"tinvest-go/internal/pkg/utils"
	tinvest_service "tinvest-go/internal/service/tinvest"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *api) HandleAccountPosition(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccountPosition.FindAllStringSubmatch(request.Caption, -1)
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

	if isCommand(request.Text) {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.AccountPositionHelp, account.Name))
		if err != nil {
			return err
		}

		return nil
	}

	instruments, err := a.tinvestService.GetInstruments(ctx, user.Token, accountID, nil)
	if err != nil {
		return fmt.Errorf("failed to get instrument for ticker %s: %+v", request.Text, err)
	}

	instrumentsByTicker := utils.ToMap(instruments, func(instrument *model.Instrument) (string, *model.Instrument) {
		return instrument.Ticker, instrument
	})

	instrument, exists := instrumentsByTicker[strings.ToUpper(request.Text)]
	if !exists {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.InstrumentNotFound, request.Text))
		if err != nil {
			return err
		}
		return nil
	}

	positions, err := a.tinvestService.GetPositions(ctx, user.Token, accountID, tinvest_service.From, time.Now(), nil)
	if err != nil {
		return fmt.Errorf("failed to get position for ticker %s: %+v", instrument.Ticker, err)
	}

	positionsByTicker := utils.ToMap(positions, func(position *model.Position) (string, *model.Position) {
		return position.Ticker, position
	})

	position, exists := positionsByTicker[instrument.Ticker]
	if !exists {
		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, fmt.Sprintf(texts.InstrumentNotFound, request.Text))
		if err != nil {
			return err
		}
		return nil
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(texts.AccountPositionDetailTitle, position.Ticker, account.Name, time.Now().Format("02.01.2006 15:04")))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionCurrency, strings.ToUpper(position.Currency)))
	sb.WriteString("\n")

	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionQuantityBuy, position.QuantityBuy))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionPriceBuy, a.accounting.FormatMoney(position.PriceBuy), a.accounting.FormatMoney(position.PriceBuyRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionValueBuy, a.accounting.FormatMoney(position.ValueBuy), a.accounting.FormatMoney(position.ValueBuyRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionNKDBuy, a.accounting.FormatMoney(position.NKDBuy), a.accounting.FormatMoney(position.NKDBuyRub)))
	sb.WriteString("\n")

	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionQuantitySell, position.QuantitySell))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionPriceSell, a.accounting.FormatMoney(position.PriceSell), a.accounting.FormatMoney(position.PriceSellRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionValueSell, a.accounting.FormatMoney(position.ValueSell), a.accounting.FormatMoney(position.ValueSellRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionNKDSell, a.accounting.FormatMoney(position.NKDSell), a.accounting.FormatMoney(position.NKDSellRub)))
	sb.WriteString("\n")

	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionQuantityEnd, position.QuantityEnd))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionPriceEnd, a.accounting.FormatMoney(position.PriceEnd), a.accounting.FormatMoney(position.PriceEndRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionValueEnd, a.accounting.FormatMoney(position.ValueEnd), a.accounting.FormatMoney(position.ValueEndRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionNKDEnd, a.accounting.FormatMoney(position.NKDEnd), a.accounting.FormatMoney(position.NKDEndRub)))
	sb.WriteString("\n")

	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionDividends, a.accounting.FormatMoney(position.Dividends), a.accounting.FormatMoney(position.DividendsRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionCoupons, a.accounting.FormatMoney(position.Coupons), a.accounting.FormatMoney(position.CouponsRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionOvernight, a.accounting.FormatMoney(position.Overnight), a.accounting.FormatMoney(position.OvernightRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionTaxes, a.accounting.FormatMoney(position.Taxes), a.accounting.FormatMoney(position.TaxesRub)))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionCommissions, a.accounting.FormatMoney(position.Commissions), a.accounting.FormatMoney(position.CommissionsRub)))
	sb.WriteString("\n")

	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(texts.PositionTotal, a.accounting.FormatMoney(position.Total), a.accounting.FormatMoney(position.TotalRub)))
	sb.WriteString("\n")

	sb.WriteString("\n")

	_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, sb.String())
	if err != nil {
		return err
	}

	return nil
}
