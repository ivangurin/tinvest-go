package bot_client

import (
	"context"
	"fmt"
	"tinvest-go/internal/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IClient interface {
	GetUpdatesChan() tgbotapi.UpdatesChannel
	SendCallback(ctx context.Context, callbackQueryID string) error
	SendMessage(ctx context.Context, messageConfig *tgbotapi.MessageConfig) (int, error)
	SendMessageWithText(ctx context.Context, chatID int64, text string) (int, error)
	DeleteMessage(ctx context.Context, chatID int64, messageID int) error
	Close() error
}

type client struct {
	API *tgbotapi.BotAPI
}

func NewClient(ctx context.Context, token string, debug bool) IClient {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logger.Fatalf(ctx, "failed to create bot client: %s", err.Error())
	}
	api.Debug = debug

	return &client{
		API: api,
	}
}

func (c *client) GetUpdatesChan() tgbotapi.UpdatesChannel {
	return c.API.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 360,
	})
}

func (c *client) Close() error {
	c.API.StopReceivingUpdates()
	logger.Info(context.Background(), "bot client has been stopped")
	return nil
}

func (c *client) SendCallback(ctx context.Context, callbackQueryID string) error {
	callback := tgbotapi.NewCallback(callbackQueryID, "")
	_, err := c.API.Request(callback)
	if err != nil {
		return fmt.Errorf("failed to send callback: %w", err)
	}
	return nil
}

func (c *client) SendMessage(ctx context.Context, messageConfig *tgbotapi.MessageConfig) (int, error) {
	messageConfig.ParseMode = tgbotapi.ModeHTML
	message, err := c.API.Send(messageConfig)
	if err != nil {
		return 0, fmt.Errorf("failed to send message: %w", err)
	}
	return message.MessageID, nil
}

func (c *client) SendMessageWithText(ctx context.Context, chatID int64, text string) (int, error) {
	message := tgbotapi.NewMessage(chatID, text)
	return c.SendMessage(ctx, &message)
}

func (c *client) DeleteMessage(ctx context.Context, chatID int64, messageID int) error {
	deleteMessage := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := c.API.Request(deleteMessage)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	return nil
}
