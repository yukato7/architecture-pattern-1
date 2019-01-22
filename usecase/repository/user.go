package repository

import (
	"context"
	"github.com/yutify/gengo-api/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userID string) (*model.User, error)
	UserExists(ctx context.Context, userID string) (bool, error)
}
