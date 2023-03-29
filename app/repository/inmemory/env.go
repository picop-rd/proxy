package inmemory

import (
	"context"
	"sync"

	"github.com/picop-rd/proxy/app/entity"
	"github.com/picop-rd/proxy/app/repository"
)

type Env struct {
	db map[string]entity.Env
	mu sync.RWMutex
}

var _ repository.Env = &Env{}

func NewEnv() *Env {
	return &Env{
		db: map[string]entity.Env{},
	}
}

func (e *Env) Get(_ context.Context, id string) (entity.Env, error) {
	e.mu.RLock()
	env, ok := e.db[id]
	e.mu.RUnlock()
	if !ok {
		return entity.Env{}, entity.ErrNotFound
	}
	if err := env.Validate(); err != nil {
		return entity.Env{}, entity.ErrNotFound
	}
	return env, nil
}

func (e *Env) Upsert(_ context.Context, envs []entity.Env) error {
	for _, v := range envs {
		e.mu.Lock()
		e.db[v.EnvID] = v
		e.mu.Unlock()
	}
	return nil
}

func (e *Env) Delete(_ context.Context, id string) error {
	e.mu.Lock()
	e.db[id] = entity.Env{}
	e.mu.Unlock()
	return nil
}
