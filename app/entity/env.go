package entity

import "github.com/rs/zerolog"

type Env struct {
	EnvID       string `json:"env_id"`
	Destination string `json:"destination"`
}

func (e *Env) Validate() error {
	if len(e.EnvID) == 0 || len(e.Destination) == 0 {
		return ErrInvalid
	}
	return nil
}

func (e Env) MarshalZerologObject(event *zerolog.Event) {
	event.Str("EnvID", e.EnvID).
		Str("Destination", e.Destination)
}
