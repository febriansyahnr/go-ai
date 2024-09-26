package main

import (
	"fmt"
	"log"

	"github.com/febriansyahnr/go-ai/config"
	"github.com/febriansyahnr/go-ai/internal/server"
)

func main() {
	conf, secret, err := config.LoadConfig("./.config.yaml", "./.secret.yaml")
	if err != nil {
		log.Fatalf("Error Loading config files: %v", err)
		return
	}

	srv := server.New(conf, secret)
	err = srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
