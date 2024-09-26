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
			Tools: []openai.Tool{
				{
					Type: "function",
					Function: &openai.FunctionDefinition{
						Name:        "get_current_weather",
						Description: "Get the current weather in a given location",
						Parameters: map[string]any{
							"type": "object",
							"properties": map[string]any{
								"location": map[string]string{
									"type":        "string",
									"description": "The city and state, e.g. San Francisco, CA",
								},
							},
							"required": []string{"location"},
						},
					},
				},
			},
		},
	)
	if err != nil {
		return "error", err
	}
	content := resp.Choices[0].Message.Content
	if len(resp.Choices[0].Message.ToolCalls) > 0 {
		for _, tool := range resp.Choices[0].Message.ToolCalls {
			fmt.Printf("%#v\n", tool)
		}
	}
	*messages = append(*messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	return content, nil
}
