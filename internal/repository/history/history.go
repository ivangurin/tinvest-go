package history_repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/db"
	"tinvest-go/internal/pkg/logger"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type IRepository interface {
	CreateRecord(ctx context.Context, userID int64, command string) (string, error)
	GetRecords(ctx context.Context, id string) (*model.History, error)
}

type Repository struct {
	client db.IClient
}

const (
	tableHistories = "histories"

	fieldID        = "id"
	fieldUserID    = "user_id"
	fieldCommand   = "command"
	fieldCreatedAt = "created_at"
)

func NewRepository(client db.IClient) IRepository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) CreateRecord(ctx context.Context, userID int64, command string) (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		logger.Errorf(ctx, "failed to generate uuid: %s", err.Error())
	}

	history :=
		&History{
			ID:        uuid.String(),
			CreatedAt: sql.NullTime{Valid: true, Time: time.Now().UTC()},
			UserID:    sql.NullInt64{Valid: true, Int64: userID},
			Command:   sql.NullString{Valid: true, String: command},
		}

	builder := sq.
		Insert(tableHistories).
		Values(history.ID, history.CreatedAt, history.UserID, history.Command)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.client.GetDB().ExecContext(ctx, query, args...)
	if err != nil {
		return "", fmt.Errorf("failed to insert history record for user %d command %s: %w", userID, command, err)
	}

	return uuid.String(), nil
}

func (r *Repository) GetRecords(ctx context.Context, id string) (*model.History, error) {
	builder := sq.
		Select(fieldID, fieldUserID, fieldCommand, fieldCreatedAt).
		From(tableHistories).
		Where(
			sq.Eq{fieldID: id},
		).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	dbHistory := &History{}
	err = r.client.GetDB().
		QueryRowContext(ctx, query, args...).
		Scan(
			&dbHistory.ID,
			&dbHistory.UserID,
			&dbHistory.Command,
			&dbHistory.CreatedAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to select history record by id %s: %w", id, err)
	}

	return &model.History{
		ID:        dbHistory.ID,
		UserID:    dbHistory.UserID.Int64,
		Command:   dbHistory.Command.String,
		CreatedAt: dbHistory.CreatedAt.Time,
	}, nil
}
