package bot

import (
	"context"
	"regexp"
	"tinvest-go/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler struct {
	Pattern *regexp.Regexp
	Handler func(ctx context.Context, user *model.User, request *tgbotapi.Message) error
}

const (
	commandToken = "/token"
)
