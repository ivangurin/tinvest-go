package suite_provider

import (
	"context"
	exchange_service "tinvest-go/internal/service/exchange"
)

type services struct {
	exchangeService exchange_service.IService
}

func (sp *suiteProvider) GetExchangeService(ctx context.Context) exchange_service.IService {
	if sp.services.exchangeService == nil {
		sp.services.exchangeService = exchange_service.NewService(
			sp.GetCurrencyRepo(ctx),
			sp.GetCbrfClient(),
		)
	}
	return sp.services.exchangeService
}
