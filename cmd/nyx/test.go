package nyx

import (
	"log"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
    Use: "test",
    Aliases: []string{"te"},
    Short: "Simple test command",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string){
        log.Printf("Hello from test: %s", args[0])
    },
}

func init() {
    rootCmd.AddCommand(testCmd)
}
