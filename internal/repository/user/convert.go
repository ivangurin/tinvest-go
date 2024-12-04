package user_repo

import (
	"database/sql"
	"tinvest-go/internal/model"
	"tinvest-go/internal/pkg/utils"
)

func toModel(user *User) *model.User {
	return &model.User{
		ID:        user.ID,
		ChatID:    user.ChatID.Int64,
		Username:  user.Username.String,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      user.Role.String,
		Token:     user.Token.String,
		CreatedAt: utils.IfThenElse(user.CreatedAt.Valid, &user.CreatedAt.Time, nil),
		ChangedAt: utils.IfThenElse(user.ChangedAt.Valid, &user.ChangedAt.Time, nil),
		DeletedAt: utils.IfThenElse(user.DeletedAt.Valid, &user.DeletedAt.Time, nil),
	}
}

func toDbModel(user *model.User) *User {
	createdAt := sql.NullTime{}
	if user.CreatedAt != nil {
		createdAt.Valid = true
		createdAt.Time = *user.CreatedAt
	}

	changedAt := sql.NullTime{}
	if user.ChangedAt != nil {
		changedAt.Valid = true
		changedAt.Time = *user.ChangedAt
	}

	deletedAt := sql.NullTime{}
	if user.DeletedAt != nil {
		deletedAt.Valid = true
		deletedAt.Time = *user.DeletedAt
	}

	return &User{
		ID:        user.ID,
		ChatID:    sql.NullInt64{Valid: user.ChatID != 0, Int64: user.ChatID},
		Username:  sql.NullString{Valid: user.Username != "", String: user.Username},
		FirstName: sql.NullString{Valid: user.FirstName != "", String: user.FirstName},
		LastName:  sql.NullString{Valid: user.LastName != "", String: user.LastName},
		Role:      sql.NullString{Valid: user.Role != "", String: user.Role},
		Token:     sql.NullString{Valid: user.Token != "", String: user.Token},
		CreatedAt: createdAt,
		ChangedAt: changedAt,
		DeletedAt: deletedAt,
	}
}
