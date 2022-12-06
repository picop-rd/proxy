package main

import (
	"flag"

	"github.com/hiroyaonoe/bcop-proxy/admin/api/http"
	"github.com/hiroyaonoe/bcop-proxy/admin/api/http/controller"
	"github.com/hiroyaonoe/bcop-proxy/admin/usecase"
	"github.com/hiroyaonoe/bcop-proxy/repository/mysql"
	"github.com/rs/zerolog/log"
)

func main() {
	dsn := flag.String("mysql", "", "mysql data source name")
	address := flag.String("port", ":8080", "server address")
	flag.Parse()

	db, err := mysql.NewDB(*dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect mysql server")
	}
	defer db.Close()

	repoEnv := mysql.NewEnv(db)
	ucEnv := usecase.NewEnv(repoEnv)
	ctrlEnv := controller.NewEnv(ucEnv)
	router := http.NewRouter(ctrlEnv)

	router.Run(*address)
}
