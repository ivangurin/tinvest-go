package service_provider

import (
	"context"
	"time"
	"tinvest-go/internal/config"
	bot_client "tinvest-go/internal/pkg/client/bot"
	cbrf_client "tinvest-go/internal/pkg/client/cbrf"
	tinvest_client "tinvest-go/internal/pkg/client/tinvest"
	"tinvest-go/internal/pkg/grpc"
)

type clients struct {
	cbrfClient    cbrf_client.IClient
	tinvestClient tinvest_client.IClient
	botClient     bot_client.IClient
}

func (sp *ServiceProvider) GetCbrfClient(ctx context.Context) cbrf_client.IClient {
	if sp.clients.cbrfClient == nil {
		sp.clients.cbrfClient = cbrf_client.NewClient()
	}
	return sp.clients.cbrfClient
}

func (sp *ServiceProvider) GetTinvestClient() tinvest_client.IClient {
	if sp.clients.tinvestClient == nil {
		sp.clients.tinvestClient = tinvest_client.NewClient(
			grpc.NewConnection(
				config.AppName,
				tinvest_client.ProdHost,
				time.Minute),
		)
	}
	return sp.clients.tinvestClient
}

func (sp *ServiceProvider) GetBotClient(ctx context.Context) bot_client.IClient {
	if sp.clients.botClient == nil {
		sp.clients.botClient = bot_client.NewClient(ctx, config.BotToken, config.BotDebug)
		sp.GetCloser().Add(sp.clients.botClient.Close)
	}
	return sp.clients.botClient
}
