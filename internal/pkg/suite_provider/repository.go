package suite_provider

import (
	"context"

	"tinvest-go/internal/config"
	"tinvest-go/internal/pkg/db"
	"tinvest-go/internal/pkg/logger"
	currency_repo "tinvest-go/internal/repository/currency"
	history_repo "tinvest-go/internal/repository/history"
	user_repo "tinvest-go/internal/repository/user"
)

type repositories struct {
	dbClient     db.IClient
	userRepo     user_repo.IRepository
	currencyRepo currency_repo.IRepository
	historyRepo  history_repo.IRepository
}

func (sp *suiteProvider) GetDBClient(ctx context.Context) db.IClient {
	if sp.repositories.dbClient == nil {
		client, err := db.NewClient(config.DbDsnTest)
		if err != nil {
			logger.Fatalf(ctx, "failed to create db client: %s", err.Error())
		}
		sp.repositories.dbClient = client
		sp.GetCloser().Add(client.Close)
	}
	return sp.repositories.dbClient
}

func (sp *suiteProvider) GetUserRepo(ctx context.Context) user_repo.IRepository {
	if sp.repositories.userRepo == nil {
		sp.repositories.userRepo = user_repo.NewRepository(
			sp.GetDBClient(ctx),
		)
	}
	return sp.repositories.userRepo
}

func (sp *suiteProvider) GetCurrencyRepo(ctx context.Context) currency_repo.IRepository {
	if sp.repositories.currencyRepo == nil {
		sp.repositories.currencyRepo = currency_repo.NewRepository(
			sp.GetDBClient(ctx),
		)
	}
	return sp.repositories.currencyRepo
}

func (sp *suiteProvider) GetHistoryRepo(ctx context.Context) history_repo.IRepository {
	if sp.repositories.historyRepo == nil {
		sp.repositories.historyRepo = history_repo.NewRepository(
			sp.GetDBClient(ctx),
		)
	}
	return sp.repositories.historyRepo
}
