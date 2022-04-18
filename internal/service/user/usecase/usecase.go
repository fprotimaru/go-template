package userusecase

import (
	"context"

	"project/internal/entity"
	"project/internal/service/user"
)

type UseCase struct {
	repo userservice.Repository
}

func NewUseCase(repo userservice.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (u UseCase) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.repo.Create(ctx, user)
}
