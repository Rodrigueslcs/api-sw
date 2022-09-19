package server

import (
	"api-sw/src/core/api/handlers"
	"api-sw/src/core/api/routes"
	"api-sw/src/server/config"
	"api-sw/src/shared/middlewares"
	"api-sw/src/shared/providers/logger"
	"fmt"
	"net/http"
	"os"
)

type server struct {
	Addr       string
	Port       int
	httpServer http.Server
}

func New() *server {
	return &server{
		Addr: "0.0.0.0",
		Port: 8081,
	}
}

func (s server) Run(cfg config.Config, log logger.ILoggerProvider) error {
	log.Info("server.main.Run", fmt.Sprintf("Server running on port : %d", s.Port))
	log.Info("server.main.Run", fmt.Sprintf("Environment : %s", cfg.Environment))
	log.Info("server.main.Run", fmt.Sprintf("Version : %s", cfg.Version))

	HandlerDependencies := handlers.Dependencies{Logger: log}
	router := routes.NewRoutes(handlers.NewHandler(HandlerDependencies))
	router.Setup()

	s.httpServer = http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Addr, s.Port),
		Handler: middlewares.Recovery(router.Client),
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
			}

		} else {
			fmt.Println(err)
		}
		process, err := os.FindProcess(os.Getegid())

		if err != nil {
			fmt.Println("couldn't find process to exit")
			os.Exit(1)
		}
		if err := process.Signal(os.Interrupt); err != nil {
			fmt.Println("couldn't find process to exit")
			os.Exit(1)
		}
	}()

	return nil
}
