package openapi

import (
	"fmt"
	"os"

	"github.com/TravisBubb/nyx-core/pkg/openapi"
	"github.com/spf13/cobra"
)

var validateOpenAPIV3Cmd = &cobra.Command{
    Use: "validate",
    Short: "Validate the given OpenAPI specification file",
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error{
        bytes, err := os.ReadFile(args[0])
        if err != nil{
            return err
        }

        validator := &openapi.OpenAPIv3Validator{}
        err = validator.Validate(bytes)
        
        if err != nil{
            return err
        }
        
        fmt.Fprintf(cmd.OutOrStdout(), "[OK] %s is a valid OpenAPI V3 specification\n", args[0])
        return nil
    },
}

func init() {
    openapiCmd.AddCommand(validateOpenAPIV3Cmd)
}
