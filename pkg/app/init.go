package app

import (
	"net/http"
	"handsongolang/pkg/router"
	"handsongolang/pkg/server"
	"handsongolang/pkg/config"
	"handsongolang/pkg/user"
	"handsongolang/pkg/db"
	"io"
	"os"
	"handsongolang/pkg/reporters"
	"go.uber.org/zap"
	"log"
)

func initHTTPServer(configFile string) {
    cfg := config.NewConfig(configFile)
    lgr := initLogger(cfg)
	rt := initRouter(cfg)

	server.NewServer(cfg, lgr ,rt).Start()
}

func initRouter(cfg config.Config) http.Handler {
	userRepo := initRepository(cfg)
	userService := initService(userRepo)
    return router.Routers(userService)
}

func initService(userRepo user.UserRepository) (user.Service){
	userService := user.NewUserService(userRepo)

	return userService
}

func initRepository(cfg config.Config) (user.UserRepository) {
	dbConfig := cfg.GetDBConfig()
	dbHandler := db.NewDBHandler(dbConfig)

	db, err := dbHandler.GetDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	return user.NewUserRepository(db)	
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



