package mysql

import (
	"github.com/hiroyaonoe/bcop-proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/repository"
)

type Env struct {
	db *DB
}

var _ repository.Env = &Env{}

func NewEnv(db *DB) *Env {
	return &Env{db: db}
}

func (e *Env) Get(id string) (entity.Env, error) {
	return entity.Env{}, nil
}

func (e *Env) Upsert(env entity.Env) error {
	return nil
}

func (e *Env) Delete(id string) error {
	return nil
}
