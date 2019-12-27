package service

import (
	"context"
	"gophr/api/v1/user"
	"gophr/model"
)

func New(repo user.Repository) *Service {
	return &Service{ repo: repo }
}

type Service struct {
	repo user.Repository
}

func(s *Service) GetByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func(s *Service) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func(s *Service) GetByUsername(ctx context.Context, uname string) (*model.User, error) {
	return s.repo.GetByUsername(ctx, uname)
}

func(s *Service) Save(ctx context.Context, usr *model.User) error {
	return s.repo.Save(ctx, usr)
}
