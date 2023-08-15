package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"github.com/zacharyrgonzales-portfolio/ChatBridge/cmd"
)

// Root Command
var rootCmd = &cobra.Command{Use: "chatbridge"}


// Main
func main() {
	rootCmd.AddCommand(cmd.startConversationCmd) // Add the subcommands here

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}