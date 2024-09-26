package ai

import (
	"github.com/febriansyahnr/go-ai/config"
	"github.com/febriansyahnr/go-ai/internal/service"
	"github.com/sashabaranov/go-openai"
)

type AI struct {
	client *openai.Client
	Config *config.Config
	Secret *config.Secret
}

func New(config *config.Config, secret *config.Secret, client *openai.Client) *AI {
	ai := &AI{
		client: client,
		Config: config,
		Secret: secret,
	}
	return ai
}

var _ service.IAI = (*AI)(nil)
