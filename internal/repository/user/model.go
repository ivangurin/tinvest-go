package user_repo

import "database/sql"

type User struct {
	ID        int64          `db:"id"`
	ChatID    sql.NullInt64  `db:"chat_id"`
	Username  sql.NullString `db:"username"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Role      sql.NullString `db:"role"`
	Token     sql.NullString `db:"token"`
	CreatedAt sql.NullTime   `db:"created_at"`
	ChangedAt sql.NullTime   `db:"changed_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}
