
package run 

import (
	"github.com/TravisBubb/nyx-core/cmd/nyx"
	"github.com/spf13/cobra"
)


var runCmd = &cobra.Command{
    Use: "run",
    Short: "Perform various operations around running API tests",
}

func init() {
    nyx.RootCmd.AddCommand(runCmd)
}
