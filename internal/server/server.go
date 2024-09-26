package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/febriansyahnr/go-ai/config"
)

type Server struct {
	port int
}

func New(conf *config.Config, secret *config.Secret) *http.Server {
	NewServer := &Server{
		port: conf.Port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
