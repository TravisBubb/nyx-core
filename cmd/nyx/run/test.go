package run

import (
	"fmt"

	"github.com/TravisBubb/nyx-core/internal/services/parser"
	"github.com/TravisBubb/nyx-core/internal/services/runner"
	"github.com/spf13/cobra"
)

var runTestCmd = &cobra.Command{
	Use:   "test <test_file>",
	Short: "Execute the provided test suite",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p := parser.NewJsonTestParser()
		ts, err := p.ParseFile(args[0])
		if err != nil {
			return err
		}

		r := runner.NewTestRunner()
		res, err := r.Execute(ts)
		if err != nil {
			return err
		}

		// TODO: process the test results and print a summary for the user
		// create a separate service for printing the results in a nice format.
		// Could I maybe use something like BubbleTea or Lipgloss for a "pretty"
		// CLI?

		fmt.Println("\n----------------")
		fmt.Printf("Requests Sent: %d\nTests Passed: %d\nTests Failed: %d\n",
			res.RequestCount, res.TestsPassed, res.TestsFailed)

		return nil
	},
}

func init() {
	runCmd.AddCommand(runTestCmd)
}
