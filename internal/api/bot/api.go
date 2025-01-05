package bot

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	bot_client "tinvest-go/internal/pkg/client/bot"
	"tinvest-go/internal/pkg/logger"
	history_service "tinvest-go/internal/service/history"
	tinvest_service "tinvest-go/internal/service/tinvest"
	user_service "tinvest-go/internal/service/user"
	"tinvest-go/internal/texts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/leekchan/accounting"
)

type API interface {
	Serve(ctx context.Context) error
}

type api struct {
	botClient      bot_client.IClient
	userService    user_service.IService
	tinvestService tinvest_service.IService
	historyService history_service.IService
	accounting     accounting.Accounting
	handlers       []handler
	lastCommand    map[int64]string
}

const (
	commandAccount          = "/account/%s"
	commandAccountTotals    = "/account/%s/totals"
	commandAccountPortfolio = "/account/%s/portfolio"
	commandAccountDetail    = "/account/%s/detail"
	commandAccountPositions = "/account/%s/positions"
	commandAccountPosition  = "/account/%s/position"
	commandAccountTrades    = "/account/%s/trades"
	commandAccountTradesFor = "/account/%s/trades-for/%s"

	tradesForCurrDay   = "currDay"
	tradesForPrevDay   = "prevDay"
	tradesForCurrWeek  = "currWeek"
	tradesForPrevWeek  = "prevWeek"
	tradesForCurrMonth = "currMonth"
	tradesForPrevMonth = "prevMonth"
)

var (
	regexpStart            *regexp.Regexp = newRegexp(`^\/start$`)
	regexpToken            *regexp.Regexp = newRegexp(`^\/token$`)
	regexpAccounts         *regexp.Regexp = newRegexp(`^\/accounts$`)
	regexpAccount          *regexp.Regexp = newRegexp(`^\/account\/([0-9]+)$`)
	regexpAccountTotals    *regexp.Regexp = newRegexp(`^\/account\/(.+)\/totals$`)
	regexpAccountPortfolio *regexp.Regexp = newRegexp(`^\/account\/(.+)\/portfolio$`)
	regexpAccountDetail    *regexp.Regexp = newRegexp(`^\/account\/(.+)\/detail$`)
	regexpAccountPositions *regexp.Regexp = newRegexp(`^\/account\/(.+)\/positions$`)
	regexpAccountPosition  *regexp.Regexp = newRegexp(`^\/account\/(.+)\/position$`)
	regexpAccountTrades    *regexp.Regexp = newRegexp(`^\/account\/(.+)\/trades$`)
	regexpAccountTradesFor *regexp.Regexp = newRegexp(`^\/account\/(.+)\/trades-for\/(.+)$`)
	regexpRsiDaily         *regexp.Regexp = newRegexp(`^\/rsi$`)
)

func NewAPI(
	botClient bot_client.IClient,
	userService user_service.IService,
	tinvestService tinvest_service.IService,
	historyService history_service.IService,
) API {
	a := &api{
		botClient:      botClient,
		userService:    userService,
		tinvestService: tinvestService,
		historyService: historyService,
		accounting: accounting.Accounting{
			Precision: 2,
			Thousand:  ".",
			Decimal:   ",",
		},
		lastCommand: map[int64]string{},
	}

	a.handlers = []handler{
		{
			Pattern: regexpStart,
			Handler: a.HandleStart,
		},
		{
			Pattern: regexpToken,
			Handler: a.HandleToken,
		},
		{
			Pattern: regexpAccounts,
			Handler: a.HandleAccounts,
		},
		{
			Pattern: regexpAccount,
			Handler: a.HandleAccount,
		},
		{
			Pattern: regexpAccountPortfolio,
			Handler: a.HandleAccountPortfolio,
		},
		{
			Pattern: regexpAccountTotals,
			Handler: a.HandleAccountTotals,
		},
		{
			Pattern: regexpAccountDetail,
			Handler: a.HandleAccountDetail,
		},
		{
			Pattern: regexpAccountPositions,
			Handler: a.HandleAccountPositions,
		},
		{
			Pattern: regexpAccountPosition,
			Handler: a.HandleAccountPosition,
		},
		{
			Pattern: regexpAccountTrades,
			Handler: a.HandleAccountTrades,
		},
		{
			Pattern: regexpAccountTradesFor,
			Handler: a.HandleAccountTradesFor,
		},
		{
			Pattern: regexpRsiDaily,
			Handler: a.HandleRSIDaily,
		},
	}

	return a
}

func (a *api) Serve(ctx context.Context) error {
	updates := a.botClient.GetUpdatesChan()

	for {
		select {
		case <-ctx.Done():
			return nil
		case update := <-updates:
			go func() {
				defer func() {
					if r := recover(); r != nil {
						logger.Errorf(ctx, "panic: %v", r)
					}
				}()

				err := a.handleRequest(ctx, &update)
				if err != nil {
					logger.Errorf(ctx, "failed to handle request: %s", err.Error())
				}
			}()
		}
	}
}

func (a *api) handleRequest(ctx context.Context, update *tgbotapi.Update) error {
	message := update.Message
	if message == nil && update.CallbackQuery != nil {
		message = update.CallbackQuery.Message
		message.From = update.CallbackQuery.From
		message.Text = update.CallbackQuery.Data
		defer func() {
			err := a.botClient.SendCallback(ctx, update.CallbackQuery.ID)
			if err != nil {
				logger.Errorf(ctx, "failed to send callback on request: %s", err.Error())
			}
		}()
	}
	if message == nil {
		return nil
	}
	if message.From.IsBot {
		return nil
	}

	_, err := a.historyService.CreateRecord(ctx, message.From.ID, message.Text)
	if err != nil {
		logger.Errorf(ctx, "failed to create history record: %s", err.Error())
	}

	user, err := a.userService.GetUserByID(ctx, message.From.ID)
	if err != nil {
		return fmt.Errorf("failed to get user %d: %w", message.From.ID, err)
	}

	var command string
	if isCommand(message.Text) {
		command = message.Text
		a.lastCommand[message.From.ID] = command
	} else {
		command = a.lastCommand[message.From.ID]
	}

	message.Caption = command

	if !regexpStart.MatchString(command) &&
		!regexpToken.MatchString(command) {
		if user == nil || user.Token == "" {
			_, err := a.botClient.SendMessageWithText(ctx, message.Chat.ID, texts.TokenIsEmpty)
			if err != nil {
				return err
			}
			return nil
		}
	}

	for _, handler := range a.handlers {
		if handler.Pattern.MatchString(command) {
			err = handler.Handler(ctx, user, message)
			if err != nil {
				return fmt.Errorf("failed to handle request %s from user %d: %w", command, message.From.ID, err)
			}
		}
	}

	return nil
}

func newRegexp(expr string) *regexp.Regexp {
	regexp, err := regexp.Compile(expr)
	if err != nil {
		panic(err.Error())
	}
	return regexp
}

func isCommand(text string) bool {
	return strings.HasPrefix(text, "/")
}
