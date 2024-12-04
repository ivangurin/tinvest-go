package model

import (
	"time"
	contractv1 "tinvest-go/internal/pb"
)

type Account struct {
	ID       string
	Type     contractv1.AccountType
	Name     string
	OpenedAt time.Time
	ClosedAt *time.Time
}

type Accounts []*Account
