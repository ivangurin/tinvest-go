package user_service

import (
	"context"
	"tinvest-go/internal/model"
	user_repo "tinvest-go/internal/repository/user"
)

type IService interface {
	IsUserExists(ctx context.Context, id int64) (bool, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type service struct {
	userRepo user_repo.IRepository
}

func NewService(userRepo user_repo.IRepository) IService {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) IsUserExists(ctx context.Context, id int64) (bool, error) {
	return s.userRepo.IsUserExists(ctx, id)
}

func (s *service) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *service) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *service) DeleteUser(ctx context.Context, id int64) error {
	return s.userRepo.DeleteUser(ctx, id)
}
