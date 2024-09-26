package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/febriansyahnr/go-ai/config"
	"github.com/febriansyahnr/go-ai/internal/service"
)

type Server struct {
	port   int
	config *config.Config
	secret *config.Secret
	AI     service.IAI
}

type ServerFunc func(*Server)

func WithAI(ai service.IAI) ServerFunc {
	return func(s *Server) {
		s.AI = ai
	}
}

func New(conf *config.Config, secret *config.Secret, confServices ...ServerFunc) *http.Server {
	NewServer := &Server{
		port:   conf.Port,
		config: conf,
		secret: secret,
	}

	for _, fn := range confServices {
		fn(NewServer)
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
