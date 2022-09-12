package server

import (
	"api-sw/src/shared/providers/logger"
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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

func (s server) Run(log logger.ILoggerProvider) error {
	log.Info("server.main.Run", fmt.Sprintf("Server running on port : %d", s.Port))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	router := chi.NewMux()
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
	})

	s.httpServer = http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Addr, s.Port),
		Handler: router,
	}

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

	return nil
}
