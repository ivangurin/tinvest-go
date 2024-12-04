package history_service

import (
	"context"
	history_repo "tinvest-go/internal/repository/history"
)

type IService interface {
	CreateRecord(ctx context.Context, userID int64, command string) (string, error)
}

type service struct {
	historyRepo history_repo.IRepository
}

func NewService(historyRepo history_repo.IRepository) IService {
	return &service{
		historyRepo: historyRepo,
	}
}

func (s *service) CreateRecord(ctx context.Context, userID int64, command string) (string, error) {
	return s.historyRepo.CreateRecord(ctx, userID, command)
}
