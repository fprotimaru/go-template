package userservice

import (
	"context"

	"smartio/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
