package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (ai *AI) SendMessage(ctx context.Context, msg string, messages *[]openai.ChatCompletionMessage) (string, error) {
	*messages = append(*messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	resp, err := ai.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4o,
			Messages: *messages,
			Tools:    []openai.Tool{},
		},
	)
	if err != nil {
		return "error", err
	}
	content := resp.Choices[0].Message.Content
	if len(resp.Choices[0].Message.ToolCalls) > 0 {
		content = "call tool(s): "
		for _, tool := range resp.Choices[0].Message.ToolCalls {
			content += fmt.Sprintf("%s ", tool.Function.Name)
		}
	}
	*messages = append(*messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	return content, nil
}
