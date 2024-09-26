package disbursement

import "github.com/sashabaranov/go-openai"

const CreateSingleDisbursement string = "CREATE_SINGLE_DISBURSEMENT"

func GetToolDefinition(fn string) openai.Tool {
	return openai.Tool{}
}
