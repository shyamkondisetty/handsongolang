package server

import (
	"context"
	"handsongolang/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"time"
	"log"
	"syscall"
)

type Server interface {
	Start()
}

type appServer struct {
	cfg         config.Config
	router      http.Handler
}

func (s *appServer) Start() {
	server := newHTTPServer(s.cfg.GetHTTPServerConfig(), s.router)
	log.Print("listening on %s", s.cfg.GetHTTPServerConfig().GetAddress())
	go func() { _ = server.ListenAndServe() }()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigCh

	err := server.Shutdown(context.Background())
	if err != nil {
		log.Print(err.Error())
		return
	}
}

func newHTTPServer(cfg config.HTTPServerConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         cfg.GetAddress(),
		WriteTimeout: time.Second * time.Duration(cfg.GetReadTimeout()),
		ReadTimeout:  time.Second * time.Duration(cfg.GetWriteTimeout()),
	}
}

func NewServer(cfg config.Config, router http.Handler ) Server {
	return &appServer{
		cfg:         cfg,
		router:      router,
	}
}
