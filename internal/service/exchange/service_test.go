package exchange_service_test

import (
	"context"
	"testing"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/suite_provider"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestExchange(t *testing.T) {
	type testCase struct {
		Name     string
		From     string
		FromRate float64
		To       string
		ToRate   float64
		Date     time.Time
		ValueSrc float64
		ValueDst float64
	}

	testCases := []testCase{
		{
			Name:     "RUB to RUB",
			From:     model.CurrencyRUB,
			FromRate: 1,
			To:       model.CurrencyRUB,
			ToRate:   1,
			Date:     time.Now().UTC().AddDate(0, 0, 1),
			ValueSrc: 1,
			ValueDst: 1,
		},
		{
			Name:     "USD to RUB",
			From:     model.CurrencyUSD,
			FromRate: 100,
			To:       model.CurrencyRUB,
			ToRate:   1,
			Date:     time.Now().UTC().AddDate(0, 0, 2),
			ValueSrc: 10,
			ValueDst: 1000,
		},
		{
			Name:     "RUB to USD",
			From:     model.CurrencyRUB,
			FromRate: 1,
			To:       model.CurrencyUSD,
			ToRate:   100,
			Date:     time.Now().UTC().AddDate(0, 0, 3),
			ValueSrc: 1000,
			ValueDst: 10,
		},
		{
			Name:     "USD to USD",
			From:     model.CurrencyUSD,
			FromRate: 100,
			To:       model.CurrencyUSD,
			ToRate:   100,
			Date:     time.Now().UTC().AddDate(0, 0, 3),
			ValueSrc: 10,
			ValueDst: 10,
		},
		{
			Name:     "EUR to RUB",
			From:     model.CurrencyEUR,
			FromRate: 110,
			To:       model.CurrencyRUB,
			ToRate:   1,
			Date:     time.Now().UTC().AddDate(0, 0, 2),
			ValueSrc: 10,
			ValueDst: 1100,
		},
		{
			Name:     "RUB to EUR",
			From:     model.CurrencyRUB,
			FromRate: 1,
			To:       model.CurrencyEUR,
			ToRate:   110,
			Date:     time.Now().UTC().AddDate(0, 0, 3),
			ValueSrc: 1100,
			ValueDst: 10,
		},
		{
			Name:     "EUR to EUR",
			From:     model.CurrencyEUR,
			FromRate: 110,
			To:       model.CurrencyEUR,
			ToRate:   110,
			Date:     time.Now().UTC().AddDate(0, 0, 3),
			ValueSrc: 10,
			ValueDst: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			sp, cancel := suite_provider.NewSuiteProvider()
			defer cancel()

			if tc.From != model.CurrencyRUB {
				sp.GetCbrfClientMock().EXPECT().
					GetExchangeRate(mock.Anything, tc.From, tc.Date).
					Return(tc.FromRate, nil)
			}

			if tc.To != model.CurrencyRUB {
				sp.GetCbrfClientMock().EXPECT().
					GetExchangeRate(mock.Anything, tc.To, tc.Date).
					Return(tc.ToRate, nil)
			}

			value, err := sp.GetExchangeService(ctx).Convert(ctx, tc.From, tc.To, tc.ValueSrc, tc.Date)
			require.NoError(t, err)
			require.Equal(t, tc.ValueDst, value)

			if tc.From != model.CurrencyRUB {
				rate, err := sp.GetCurrencyRepo(ctx).GetExchangeRate(ctx, tc.From, tc.Date)
				require.NoError(t, err)
				require.Equal(t, tc.FromRate, rate)
			}

			if tc.To != model.CurrencyRUB {
				rate, err := sp.GetCurrencyRepo(ctx).GetExchangeRate(ctx, tc.To, tc.Date)
				require.NoError(t, err)
				require.Equal(t, tc.ToRate, rate)
			}

		})
	}
}
