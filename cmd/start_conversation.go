package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var StartConversationCmd = &cobra.Command{
	Use:   "start-conversation",
	Short: "Start a new conversation with ChatGPT",
	Run: func(cmd *cobra.Command, args []string) {
		instruction, _ := cmd.Flags().GetString("instruction")
		response, err := startConversation(instruction)
		if err != nil {
			fmt.Println("Error starting conversation:", err)
			return
		}
		fmt.Println("Conversation started:", response)
	},
}

func init() {
	StartConversationCmd.Flags().StringP("instruction", "i", "", "Custom instruction for the conversation")
}

func startConversation(instruction string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}