package bot

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/logger"
	tinvest_service "tinvest-go/internal/service/tinvest"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const positionsPerMessage = 50

func (a *api) HandleAccountPositions(ctx context.Context, user *model.User, request *tgbotapi.Message) error {
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

	result := regexpAccountPositions.FindAllStringSubmatch(request.Text, -1)
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

	positions, err := a.tinvestService.GetPositions(ctx, user.Token, accountID, tinvest_service.From, time.Now(), nil)
	if err != nil {
		return err
	}

	messageCounter, rest := math.Modf(float64(len(positions)) / float64(positionsPerMessage))
	if rest > 0 {
		messageCounter += 1
	}

	var sb strings.Builder
	var fromPosition int
	var toPosition int
	for i := 0; i < int(messageCounter); i++ {
		fromPosition = i * positionsPerMessage
		toPosition = fromPosition + positionsPerMessage
		if toPosition > len(positions) {
			toPosition = len(positions)
		}

		sb = strings.Builder{}
		if i == 0 {
			sb.WriteString(fmt.Sprintf(texts.AccountPositionsTitle, account.Name, time.Now().Format("02.01.2006 15:04")))
			sb.WriteString("\n")
		}

		if messageCounter > 1 {
			sb.WriteString(fmt.Sprintf(texts.AccountPositionsPart, i+1, int(messageCounter)))
			sb.WriteString("\n")
		}

		for _, position := range positions[fromPosition:toPosition] {
			if position.Ticker == "" {
				continue
			}

			sb.WriteString(
				fmt.Sprintf(texts.PositionResult,
					position.Ticker,
					a.accounting.FormatMoney(position.Total),
					position.Currency,
					a.accounting.FormatMoney(position.TotalRub)))

		}

		_, err = a.botClient.SendMessageWithText(ctx, request.Chat.ID, sb.String())
		if err != nil {
			return err
		}
	}

	return nil
}
