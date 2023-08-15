package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zacharyrgonzales-portfolio/ChatBridge/cmd"
	"os"
)

// Root Command
var rootCmd = &cobra.Command{Use: "chatbridge"}

// Main
func main() {
	rootCmd.AddCommand(cmd.StartConversationCmd) // Add the start conversation command

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
