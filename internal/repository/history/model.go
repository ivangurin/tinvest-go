package history_repo

import "database/sql"

type History struct {
	ID        string         `db:"id"`
	UserID    sql.NullInt64  `db:"user_id"`
	Command   sql.NullString `db:"command"`
	CreatedAt sql.NullTime   `db:"created_at"`
}
