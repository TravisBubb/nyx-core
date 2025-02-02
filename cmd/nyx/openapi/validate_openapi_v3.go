package openapi 

import (
	"log"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
    Use: "validate",
    Short: "Validate the given OpenAPI specification file",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string){
        log.Printf("Hello from test: %s", args[0])
    },
}

func init() {
    openapiCmd.AddCommand(validateCmd)
}
