package currency_repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"tinvest-go/internal/pkg/db"

	sq "github.com/Masterminds/squirrel"
)

type IRepository interface {
	AddExchangeRate(ctx context.Context, currencyID string, date time.Time, rate float64) error
	GetExchangeRate(ctx context.Context, currencyID string, date time.Time) (float64, error)
}

type Repository struct {
	client db.IClient
}

const (
	tableCurrencies = "currencies"

	fieldID        = "id"
	fieldDate      = "date"
	fieldRate      = "rate"
	fieldChangedAt = "changed_at"
)

func NewRepository(client db.IClient) IRepository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) AddExchangeRate(ctx context.Context, currencyID string, date time.Time, rate float64) error {
	onDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	currency := &Currency{
		ID:        currencyID,
		Date:      onDate,
		Rate:      sql.NullFloat64{Valid: true, Float64: rate},
		ChangedAt: sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	builder := sq.
		Insert(tableCurrencies).
		Values(currency.ID, currency.Date, currency.Rate, currency.ChangedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.client.GetDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert currency %s on date %v: %w", currencyID, onDate, err)
	}

	return nil
}

func (r *Repository) GetExchangeRate(ctx context.Context, currencyID string, date time.Time) (float64, error) {
	onDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	builder := sq.Select(fieldRate).
		From(tableCurrencies).
		Where(
			sq.And{
				sq.Eq{fieldID: currencyID},
				sq.Eq{fieldDate: onDate},
			},
		).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	var currency sql.NullFloat64
	err = r.client.GetDB().QueryRowContext(ctx, query, args...).Scan(&currency)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to select currency id = %s, date = %v: %w", currencyID, onDate, err)
	}

	return currency.Float64, nil
}
