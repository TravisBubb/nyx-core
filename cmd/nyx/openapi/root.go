package openapi

import (
	"github.com/TravisBubb/nyx-core/cmd/nyx"
	"github.com/spf13/cobra"
)


var openapiCmd = &cobra.Command{
    Use: "openapi",
    Short: "Perform various operations on an open api specification",
}

func init() {
    nyx.RootCmd.AddCommand(openapiCmd)
}
