package entity

type Env struct {
	EnvID       string
	Destination string
}

func (e *Env) Validate() error {
	if len(e.EnvID) == 0 || len(e.Destination) == 0 {
		return ErrInvalid
	}
	return nil
}
