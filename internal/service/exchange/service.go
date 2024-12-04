package exchange_service

import (
	"context"
	"fmt"
	"math"
	"time"
	"tinvest-go/internal/model"
	cbrf_client "tinvest-go/internal/pkg/client/cbrf"
	currency_repo "tinvest-go/internal/repository/currency"
)

type IService interface {
	GetRate(ctx context.Context, currencyID string, date time.Time) (float64, error)
	Convert(ctx context.Context, from string, to string, value float64, date time.Time) (float64, error)
}

type service struct {
	currencyRepo currency_repo.IRepository
	cbrfClient   cbrf_client.IClient
}

func NewService(
	currencyRepo currency_repo.IRepository,
	cbrfClient cbrf_client.IClient,
) IService {
	return &service{
		currencyRepo: currencyRepo,
		cbrfClient:   cbrfClient,
	}
}

func (s *service) GetRate(ctx context.Context, currencyID string, date time.Time) (float64, error) {
	rate, err := s.currencyRepo.GetExchangeRate(ctx, currencyID, date)
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	if rate == 0 {
		rate, err = s.cbrfClient.GetExchangeRate(ctx, currencyID, date)
		if err != nil {
			return 0, fmt.Errorf("failed to get exchange rate: %w", err)
		}

		err = s.currencyRepo.AddExchangeRate(ctx, currencyID, date, rate)
		if err != nil {
			return 0, fmt.Errorf("error updating exchange rate: %s", err.Error())
		}

	}

	return rate, nil
}

func (s *service) Convert(ctx context.Context, from string, to string, value float64, date time.Time) (float64, error) {
	if from == to {
		return value, nil
	}
	if value == 0 {
		return 0, nil
	}

	result := value

	if from != model.CurrencyRUB {
		exchangeRate, err := s.GetRate(ctx, from, date)
		if err != nil {
			return 0, err
		}

		result = result * exchangeRate
	}

	if to != model.CurrencyRUB {
		exchangeRate, err := s.GetRate(ctx, to, date)
		if err != nil {
			return 0, err
		}

		result = result / exchangeRate
	}

	return (math.Floor(result*100) / 100), nil
}
