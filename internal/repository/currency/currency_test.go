package currency_repo_test

import (
	"context"
	"testing"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/suite_provider"

	"github.com/stretchr/testify/require"
)

func TestCurrency(t *testing.T) {
	type testCase struct {
		Name       string
		CurrencyID string
		Date       time.Time
		Rate       float64
	}

	testCases := []testCase{
		{
			Name:       "USD today",
			CurrencyID: model.CurrencyUSD,
			Date:       time.Now().UTC(),
			Rate:       105.33,
		},
		{
			Name:       "USD next day",
			CurrencyID: model.CurrencyUSD,
			Date:       time.Now().UTC().Add(-24 * time.Hour),
			Rate:       100.66,
		},
		{
			Name:       "USD last day",
			CurrencyID: model.CurrencyUSD,
			Date:       time.Now().UTC().Add(24 * time.Hour),
			Rate:       110.99,
		},
		{
			Name:       "EUR today",
			CurrencyID: model.CurrencyEUR,
			Date:       time.Now().UTC(),
			Rate:       115.33,
		},
		{
			Name:       "EUR last day",
			CurrencyID: model.CurrencyEUR,
			Date:       time.Now().UTC().Add(-24 * time.Hour),
			Rate:       110.66,
		},
		{
			Name:       "EUR next day",
			CurrencyID: model.CurrencyEUR,
			Date:       time.Now().UTC().Add(24 * time.Hour),
			Rate:       120.99,
		},
		{
			Name:       "CNY today",
			CurrencyID: model.CurrencyCNY,
			Date:       time.Now().UTC(),
			Rate:       14.33,
		},
		{
			Name:       "CNY last day",
			CurrencyID: model.CurrencyCNY,
			Date:       time.Now().UTC().Add(-24 * time.Hour),
			Rate:       13.66,
		},
		{
			Name:       "CNY next day",
			CurrencyID: model.CurrencyCNY,
			Date:       time.Now().UTC().Add(24 * time.Hour),
			Rate:       15.99,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			sp, cancel := suite_provider.NewSuiteProvider()
			defer cancel()

			rate, err := sp.GetCurrencyRepo(ctx).GetExchangeRate(ctx, tc.CurrencyID, tc.Date)
			require.NoError(t, err)
			require.Equal(t, 0., rate)

			err = sp.GetCurrencyRepo(ctx).AddExchangeRate(ctx, tc.CurrencyID, tc.Date, tc.Rate)
			require.NoError(t, err)

			rate, err = sp.GetCurrencyRepo(ctx).GetExchangeRate(ctx, tc.CurrencyID, tc.Date)
			require.NoError(t, err)
			require.Equal(t, tc.Rate, rate)
		})
	}
}
