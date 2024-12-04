package model

import (
	"time"
)

const (
	roleAdmin = "admin"
)

type User struct {
	ID        int64
	ChatID    int64
	Username  string
	FirstName string
	LastName  string
	Role      string
	Token     string
	CreatedAt *time.Time
	ChangedAt *time.Time
	DeletedAt *time.Time
}

func (u *User) IsAdmin() bool {
	return u.Role == roleAdmin
}
