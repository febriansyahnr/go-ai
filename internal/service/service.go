package service

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type IAI interface {
	SendMessage(ctx context.Context, msg string, messages *[]openai.ChatCompletionMessage) (string, error)
}
