package model

import "time"

type History struct {
	ID        string
	UserID    int64
	Command   string
	CreatedAt time.Time
}
