package mysql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type DB struct {
	*sqlx.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for {
		err = db.DB.Ping()
		if err == nil {
			break
		}
		log.Info().Err(err).Msg("connecting mysql server")
		time.Sleep(time.Second * 2)
	}
	log.Info().Msg("connected mysql server")

	db.DB.SetConnMaxLifetime(time.Minute * 3)
	db.DB.SetMaxIdleConns(100)
	db.DB.SetMaxOpenConns(100)
	return &DB{db}, nil
}

func (d *DB) Close() error {
	return d.DB.Close()
}
