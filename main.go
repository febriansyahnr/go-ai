package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/febriansyahnr/go-ai/config"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	_, secret, err := config.LoadConfig("./.config.yaml", "./.secret.yaml")
	if err != nil {
		log.Fatalf("Error Loading config files: %v", err)
		return
	}
	client := openai.NewClient(secret.ChatGPTToken)

	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT4o,
				Messages: messages,
			},
		)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		content := resp.Choices[0].Message.Content
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println(content)
	}
}
