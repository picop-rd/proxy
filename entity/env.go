package entity

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
