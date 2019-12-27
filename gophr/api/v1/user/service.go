package user

import (
	"context"
	"gophr/model"
)

type Service interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, id string) (*model.User, error)
	GetByUsername(ctx context.Context, uname string) (*model.User, error)
	Save(ctx context.Context, user *model.User) error
}
