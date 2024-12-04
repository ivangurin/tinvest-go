package cbrf_client

import (
	"context"
	"fmt"
	"time"

	"github.com/ivangurin/cbrf-go"
)

type IClient interface {
	GetExchangeRate(ctx context.Context, currencyID string, date time.Time) (float64, error)
}

type Client struct {
}

func NewClient() IClient {
	return &Client{}
}

func (c *Client) GetExchangeRate(ctx context.Context, currencyID string, date time.Time) (float64, error) {
	rate, err := cbrf.GetExchangeRate(ctx, currencyID, date)
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate: %w", err)
	}
	return rate, nil
}
