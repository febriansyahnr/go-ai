package disbursement

import "github.com/sashabaranov/go-openai"

const (
	CreateTopupVA            string = "CREATE_TOPUP_VA"
	CreateSingleDisbursement string = "CREATE_SINGLE_DISBURSEMENT"
)

var tools = map[string]openai.Tool{
	CreateTopupVA: {
		Type: "function",
		Function: &openai.FunctionDefinition{
			Name:        CreateTopupVA,
			Description: "create virtual account for topup",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"amount": map[string]string{
						"type":        "string",
						"description": "topup amount",
					},
				},
				"required": []string{"amount"},
			},
		},
	},
}

func GetToolDefinition(fn string) openai.Tool {
	return tools[fn]
}
