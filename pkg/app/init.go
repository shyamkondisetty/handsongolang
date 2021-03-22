package app

import (
	"net/http"
	"handsongolang/pkg/router"
	"handsongolang/pkg/server"
	"handsongolang/pkg/config"
	"io"
	"os"
	"handsongolang/pkg/reporters"
	"go.uber.org/zap"
)

func initHTTPServer(configFile string) {
    cfg := config.NewConfig(configFile)
    lgr := initLogger(cfg)
	rt := initRouter()

	server.NewServer(cfg, lgr ,rt).Start()
}

func initRouter() http.Handler {
    return router.Routers()
}

func initLogger(cfg config.Config) *zap.Logger {
	return reporters.NewLogger(
		cfg.GetLogConfig().GetLevel(),
		getWriters(cfg.GetLogFileConfig())...,
	)
}

func getWriters(cfg config.LogFileConfig) []io.Writer {
	return []io.Writer{
		os.Stdout,
		reporters.NewExternalLogFile(cfg),
	}
}



