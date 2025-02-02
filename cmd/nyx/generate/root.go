package openapi

import (
	"github.com/TravisBubb/nyx-core/cmd/nyx"
	"github.com/spf13/cobra"
)


var generateCmd = &cobra.Command{
    Use: "generate",
    Short: "Perform various operations around generating new API tests",
}

func init() {
    nyx.RootCmd.AddCommand(generateCmd)
}
