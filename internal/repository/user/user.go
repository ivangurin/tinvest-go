package user_repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/db"
	"tinvest-go/internal/pkg/utils"

	sq "github.com/Masterminds/squirrel"
)

type IRepository interface {
	IsUserExists(ctx context.Context, id int64) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
	GetUsers(ctx context.Context) ([]*model.User, error)
}

type Repository struct {
	client db.IClient
}

const (
	tableUsers     = "users"
	fieldID        = "id"
	fieldChatID    = "chat_id"
	fieldUsername  = "username"
	fieldFirstName = "first_name"
	fieldLastName  = "last_name"
	fieldRole      = "role"
	fieldToken     = "token"
	fieldCreatedAt = "created_at"
	fieldChangedAt = "changed_at"
	fieldDeletedAt = "deleted_at"
)

func NewRepository(client db.IClient) IRepository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) IsUserExists(ctx context.Context, id int64) (bool, error) {
	builder := sq.Select(fieldID).
		From(tableUsers).
		Where(
			sq.Eq{fieldID: id},
		).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to build query: %w", err)
	}

	var userID int64
	err = r.client.GetDB().QueryRowContext(ctx, query, args...).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to select user by id %d: %w", id, err)
	}

	return true, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	dbUser := toDbModel(user)
	dbUser.CreatedAt = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	builder := sq.
		Insert(tableUsers).
		Values(
			dbUser.ID,
			dbUser.ChatID,
			dbUser.Username,
			dbUser.FirstName,
			dbUser.LastName,
			dbUser.Role,
			dbUser.Token,
			dbUser.CreatedAt,
			dbUser.ChangedAt,
			dbUser.DeletedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.client.GetDB().ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user %+v: %w", user, err)
	}

	return toModel(dbUser), nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.
		Select(
			fieldID,
			fieldChatID,
			fieldUsername,
			fieldFirstName,
			fieldLastName,
			fieldRole,
			fieldToken,
			fieldCreatedAt,
			fieldChangedAt,
			fieldDeletedAt,
		).
		From(tableUsers).
		Where(
			sq.Eq{fieldID: id},
		).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	dbUser := &User{}
	err = r.client.GetDB().
		QueryRowContext(ctx, query, args...).
		Scan(
			&dbUser.ID,
			&dbUser.ChatID,
			&dbUser.Username,
			&dbUser.FirstName,
			&dbUser.LastName,
			&dbUser.Role,
			&dbUser.Token,
			&dbUser.CreatedAt,
			&dbUser.ChangedAt,
			&dbUser.DeletedAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to select user by id %d: %w", id, err)
	}

	return toModel(dbUser), nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *model.User) error {
	user.ChangedAt = utils.Ptr(time.Now().UTC())
	dbUser := toDbModel(user)

	builder := sq.
		Update(tableUsers).
		Set(fieldChatID, dbUser.ChatID).
		Set(fieldUsername, dbUser.Username).
		Set(fieldFirstName, dbUser.FirstName).
		Set(fieldLastName, dbUser.LastName).
		Set(fieldRole, dbUser.Role).
		Set(fieldToken, dbUser.Token).
		Set(fieldCreatedAt, dbUser.CreatedAt).
		Set(fieldChangedAt, dbUser.ChangedAt).
		Set(fieldDeletedAt, dbUser.DeletedAt).
		Where(
			sq.Eq{fieldID: dbUser.ID},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.client.GetDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user %+v: %w", user, err)
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int64) error {
	deletedAt := time.Now().UTC()

	builder := sq.
		Update(tableUsers).
		Set(fieldDeletedAt, deletedAt).
		Where(
			sq.Eq{fieldID: id},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.client.GetDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user %d: %w", id, err)
	}

	return nil
}

func (r *Repository) GetUsers(ctx context.Context) ([]*model.User, error) {
	builder := sq.
		Select(
			fieldID,
			fieldChatID,
			fieldUsername,
			fieldFirstName,
			fieldLastName,
			fieldRole,
			fieldToken,
			fieldCreatedAt,
			fieldChangedAt,
			fieldDeletedAt,
		).
		From(tableUsers)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := r.client.GetDB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to select users: %w", err)
	}
	defer rows.Close()

	users := []*model.User{}
	for rows.Next() {
		dbUser := &User{}
		err = rows.Scan(
			&dbUser.ID,
			&dbUser.ChatID,
			&dbUser.Username,
			&dbUser.FirstName,
			&dbUser.LastName,
			&dbUser.Role,
			&dbUser.Token,
			&dbUser.CreatedAt,
			&dbUser.ChangedAt,
			&dbUser.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, toModel(dbUser))
	}

	return users, nil
}
