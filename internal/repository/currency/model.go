package currency_repo

import (
	"database/sql"
	"time"
)

type Currency struct {
	ID        string
	Date      time.Time
	Rate      sql.NullFloat64
	ChangedAt sql.NullTime
}
