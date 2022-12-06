package main

import (
	"flag"

	"github.com/hiroyaonoe/bcop-proxy/proxy"
	"github.com/hiroyaonoe/bcop-proxy/repository/inmemory"
)

func main() {
	propagate := flag.Bool("propagate", true, "header propagation?")
	defaultAddr := flag.String("default-addr", "", "defalut address")
	port := flag.String("port", "8081", "listen port")
	flag.Parse()

	repoEnv := inmemory.NewEnv()
	server := &proxy.Server{
		Env:                         repoEnv,
		GetEnvIDFromHeaderValueFunc: proxy.GetEnvIDFromBaggage,
		Propagate:                   *propagate,
		DefaultAddr:                 *defaultAddr,
	}

	server.Start(":" + *port)
}
