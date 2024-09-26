package main

import (
	"fmt"
	"log"

	"github.com/febriansyahnr/go-ai/config"
	"github.com/febriansyahnr/go-ai/internal/server"
	"github.com/febriansyahnr/go-ai/internal/service/ai"
	"github.com/sashabaranov/go-openai"
)

func main() {
	conf, secret, err := config.LoadConfig("./.config.yaml", "./.secret.yaml")
	if err != nil {
		log.Fatalf("Error Loading config files: %v", err)
		return
	}

	chatGptAI := openai.NewClient(secret.ChatGPTToken)

	aiService := ai.New(conf, secret, chatGptAI)

	srv := server.New(conf, secret, server.WithAI(aiService))
	err = srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
