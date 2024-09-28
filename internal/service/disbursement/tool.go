package disbursement

import "github.com/sashabaranov/go-openai"

const CreateSingleDisbursement string = "CREATE_SINGLE_DISBURSEMENT"

func GetToolDefinition(fn string) openai.Tool {
	return openai.Tool{
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
	}
}
