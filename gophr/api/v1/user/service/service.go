package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gophr/api/v1"
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

func (s *Service) GetAndCompareUserPassword(ctx context.Context, username, password string) (*model.User, error) {
	// Get the users information
	usr, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// compare the password
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		// if not match then return ErrorCredential
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, v1.ErrorCredentials
		}
		return nil, err
	}

	// if match then return the user's information excluding the password
	usr.Password = ""

	return usr, nil
}