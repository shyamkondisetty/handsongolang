package app

import (
	"net/http"
	"handsongolang/pkg/router"
	"handsongolang/pkg/server"
	"handsongolang/pkg/config"
)

func initHTTPServer(configFile string) {
    cfg := config.NewConfig(configFile)
	rt := initRouter()

	server.NewServer(cfg, rt).Start()
}

func initRouter() http.Handler {
    return router.Routers()
}



