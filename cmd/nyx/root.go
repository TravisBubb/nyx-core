package nyx

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use: "nyx",
    Short: "nyx short description...",
    Long: "nyx long description...",
    Run: func (cmd *cobra.Command, args[]string) {

    },
}

func Execute() {
    if err := RootCmd.Execute(); err != nil{
        log.Printf("An error occurred during execution, %s", err)
        os.Exit(1)
    }
}
