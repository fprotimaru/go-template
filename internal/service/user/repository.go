package userservice

import (
	"context"

	"project/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
