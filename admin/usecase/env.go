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

func (e *Env) Get(ctx context.Context) ([]entity.Env, error) {
	envs, err := e.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get envs from repository: %w", err)
	}
	return envs, nil
}

func (e *Env) Register(ctx context.Context, envs []entity.Env) error {
	err := e.repo.Upsert(ctx, envs)
	if err != nil {
		return fmt.Errorf("failed to register envs to repository: %w", err)
	}
	return nil
}

func (e *Env) Delete(ctx context.Context, id string) error {
	return nil
}
