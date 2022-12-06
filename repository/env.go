package repository

import (
	"github.com/hiroyaonoe/bcop-proxy/entity"
)

type Env interface {
	Get(id string) (entity.Env, error)
	Upsert(env entity.Env) error
	Delete(id string) error
}
