package cmd

import (
	"github.com/spf13/cobra"
)

var startConversationCmd = &cobra.Command{
	Use: "start-conversation",
	Short: "Start a new conversation with ChatGPT",
	Run: func(cmd *cobra.Command, args []string) {
		// Implemtnation goes here
	},
}

func init() {
	startConversationCmd.Flags().StringP("instruction", "i", "", "Custom instruction for the conversation")
}