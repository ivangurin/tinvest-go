package service_provider

import (
	"context"
	"os"
	"syscall"
	"tinvest-go/internal/pkg/closer"
)

type ServiceProvider struct {
	ctx          context.Context
	closer       closer.Closer
	apis         apis
	clients      clients
	services     services
	repositories repositories
}

var serviceProvider *ServiceProvider

func GetServiceProvider(ctx context.Context) *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{
			ctx: ctx,
		}
	}
	return serviceProvider
}

func (sp *ServiceProvider) GetCloser() closer.Closer {
	if sp.closer == nil {
		sp.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}
	return sp.closer
}
