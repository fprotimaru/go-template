package userrepository

import (
	"context"

	"project/internal/entity"

	"github.com/uptrace/bun"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := r.NewInsert().Model(user).Exec(ctx)
	return user, err
}
