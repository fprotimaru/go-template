package userservice

import (
	"context"

	"github.com/fprotimaru/go-template/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
