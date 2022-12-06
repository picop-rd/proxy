package usecase

import (
	"context"
	"fmt"

	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	repo repository.Env
}

func NewEnv(repo repository.Env) *Env {
	return &Env{repo: repo}
}

func (e *Env) Get(ctx context.Context, id string) (entity.Env, error) {
	env, err := e.repo.Get(ctx, id)
	if err != nil {
		return entity.Env{}, fmt.Errorf("failed to get env from repository: %w", err)
	}
	return env, nil
}

func (e *Env) Register(ctx context.Context, envs []entity.Env) error {
	err := e.repo.Upsert(ctx, envs)
	if err != nil {
		return fmt.Errorf("failed to register envs to repository: %w", err)
	}
	return nil
}

func (e *Env) Delete(ctx context.Context, id string) error {
	err := e.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete env from repository: %w", err)
	}
	return nil
}
