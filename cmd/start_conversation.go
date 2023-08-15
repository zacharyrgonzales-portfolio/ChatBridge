package cmd

import (
	"github.com/spf13/cobra"
)

var StartConversationCmd = &cobra.Command{
	Use:   "start-conversation",
	Short: "Start a new conversation with ChatGPT",
	Run: func(cmd *cobra.Command, args []string) {
		// Your code to start a conversation goes here
	},
}

func init() {
	StartConversationCmd.Flags().StringP("instruction", "i", "", "Custom instruction for the conversation")
}
