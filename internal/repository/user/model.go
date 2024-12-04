package user_repo

import "database/sql"

type User struct {
	ID        int64
	ChatID    sql.NullInt64
	Username  sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
	Role      sql.NullString
	Token     sql.NullString
	CreatedAt sql.NullTime
	ChangedAt sql.NullTime
	DeletedAt sql.NullTime
}
