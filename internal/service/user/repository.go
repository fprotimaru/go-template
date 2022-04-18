package userservice

import (
	"context"

	"github.com/fprotimaru/go-template/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
