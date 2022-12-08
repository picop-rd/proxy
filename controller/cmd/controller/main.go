package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/hiroyaonoe/bcop-proxy/controller/api/http"
	"github.com/hiroyaonoe/bcop-proxy/controller/api/http/controller"
	"github.com/hiroyaonoe/bcop-proxy/controller/repository/mysql"
	"github.com/hiroyaonoe/bcop-proxy/controller/usecase"
	"github.com/rs/zerolog/log"
)

func main() {
	port := flag.String("port", "8080", "listen port")
	dsn := flag.String("mysql", "", "mysql data source name")

	flag.Parse()

	db, err := mysql.NewDB(*dsn)
	if err != nil {
		log.Fatal().Str("dsn", *dsn).Msg("failed to connect db")
	}
	defer db.Close()

	repoProxy := mysql.NewProxy(db)
	ucProxy := usecase.NewProxy(repoProxy)
	ctrlProxy := controller.NewProxy(ucProxy)

	server := http.NewServer(ctrlProxy)
	server.SetRoute()

	go server.Run(":" + *port)
	defer server.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
