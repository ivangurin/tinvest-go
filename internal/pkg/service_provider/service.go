package service_provider

import (
	"context"

	exchange_service "tinvest-go/internal/service/exchange"
	history_service "tinvest-go/internal/service/history"
	tinvest_service "tinvest-go/internal/service/tinvest"
	user_service "tinvest-go/internal/service/user"
)

type services struct {
	userService     user_service.IService
	exchangeService exchange_service.IService
	tinvestService  tinvest_service.IService
	historyService  history_service.IService
}

func (sp *ServiceProvider) GetUserService(ctx context.Context) user_service.IService {
	if sp.services.userService == nil {
		sp.services.userService = user_service.NewService(
			sp.GetUserRepo(ctx),
		)
	}
	return sp.services.userService
}

func (sp *ServiceProvider) GetExchangeService(ctx context.Context) exchange_service.IService {
	if sp.services.exchangeService == nil {
		sp.services.exchangeService = exchange_service.NewService(
			sp.GetCurrencyRepo(ctx),
			sp.GetCbrfClient(ctx),
		)
	}
	return sp.services.exchangeService
}

func (sp *ServiceProvider) GetTinvestService(ctx context.Context) tinvest_service.IService {
	if sp.services.tinvestService == nil {
		sp.services.tinvestService = tinvest_service.NewService(
			sp.GetTinvestClient(),
			sp.GetExchangeService(ctx),
		)
	}
	return sp.services.tinvestService
}

func (sp *ServiceProvider) GetHistoryService(ctx context.Context) history_service.IService {
	if sp.services.historyService == nil {
		sp.services.historyService = history_service.NewService(
			sp.GetHistoryRepo(ctx),
		)
	}
	return sp.services.historyService
}
