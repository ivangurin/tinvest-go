package bot

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"tinvest-go/internal/model"
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

var (
	regexpStart                 *regexp.Regexp = newRegexp(`^\/start$`)
	regexpToken                 *regexp.Regexp = newRegexp(`^\/token$`)
	regexpAccounts              *regexp.Regexp = newRegexp(`^\/accounts$`)
	regexpAccount               *regexp.Regexp = newRegexp(`^\/accounts\/([0-9]+)$`)
	regexpAccountTotals         *regexp.Regexp = newRegexp(`\/accounts\/(.+)\/totals`)
	regexpAccountPortfolio      *regexp.Regexp = newRegexp(`\/accounts\/(.+)\/portfolio`)
	regexpAccountDetail         *regexp.Regexp = newRegexp(`\/accounts\/(.+)\/detail`)
	regexpAccountPositions      *regexp.Regexp = newRegexp(`\/accounts\/(.+)\/positions`)
	regexpAccountPosition       *regexp.Regexp = newRegexp(`^\/accounts\/(.+)\/position$`)
	regexpAccountPositionDetail *regexp.Regexp = newRegexp(`^\/accounts\/(.+)\/position\/(.+)$`)
	regexpAccountTrades         *regexp.Regexp = newRegexp(`^\/accounts\/(.+)\/trades$`)
	regexpAccountTradesFor      *regexp.Regexp = newRegexp(`^\/accounts\/(.+)\/trades\/(.+)$`)
	regexpRsiD                  *regexp.Regexp = newRegexp(`^\/rsid$`)
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
				logger.Errorf(ctx, "api.handleRequest: %s", err.Error())
			}
		}()
	}
	if message == nil {
		return nil
	}
	if message.From.IsBot {
		return nil
	}

	user, err := a.updateUser(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	_, err = a.historyService.CreateRecord(ctx, message.From.ID, message.Text)
	if err != nil {
		return fmt.Errorf("failed to create history record: %w", err)
	}

	var command string
	if isCommand(message.Text) {
		command = message.Text
		a.lastCommand[user.ID] = command
	} else {
		command = a.lastCommand[user.ID]
	}

	for _, handler := range a.handlers {
		if handler.Pattern.MatchString(command) {
			err = handler.Handler(ctx, user, message)
			if err != nil {
				return fmt.Errorf("failed to handle command %s by user %d: %w", command, user.ID, err)
			}
		}
	}

	return nil
}

func (a *api) updateUser(ctx context.Context, message *tgbotapi.Message) (*model.User, error) {
	if message.From.IsBot {
		return nil, nil
	}

	exist, err := a.userService.IsUserExists(ctx, message.From.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existance: %w", err)
	}

	var user *model.User
	if exist {
		user, err = a.userService.GetUserByID(ctx, message.From.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get user %d: %w", user.ID, err)
		}
	} else {
		user = &model.User{
			ID:     message.From.ID,
			ChatID: message.Chat.ID,
		}
	}

	user.Username = message.From.UserName
	user.FirstName = message.From.FirstName
	user.LastName = message.From.LastName

	if exist {
		err = a.userService.UpdateUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("failed to update user %d: %w", user.ID, err)
		}
	} else {
		user, err = a.userService.CreateUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("failed to create user %d: %w", user.ID, err)
		}
	}

	return user, nil
}

func (a *api) sendMessageNoToken(ctx context.Context, user *model.User) error {
	message := tgbotapi.NewMessage(user.ChatID, texts.TokenIsEmpty)
	_, err := a.botClient.SendMessage(ctx, &message)
	if err != nil {
		return err
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
