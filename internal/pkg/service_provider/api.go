package service_provider

import (
	"context"
	"tinvest-go/internal/api/bot"
)

type apis struct {
	botApi bot.API
}

func (sp *ServiceProvider) GetBotApi(ctx context.Context) bot.API {
	if sp.apis.botApi == nil {
		sp.apis.botApi = bot.NewAPI(
			sp.GetBotClient(ctx),
			sp.GetUserService(ctx),
			sp.GetTinvestService(ctx),
			sp.GetHistoryService(ctx),
		)
	}
	return sp.apis.botApi
}
