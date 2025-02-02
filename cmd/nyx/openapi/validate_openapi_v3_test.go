package openapi

import (
	"bytes"
	"io"
	"testing"

	"github.com/TravisBubb/nyx-core/cmd/nyx"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestValidateOpenAPIV3FileNotFound(t *testing.T) {
	cmd := newRootCmd()
	cmd.SetArgs([]string{"openapi", "validate", "invalidfile.yaml"})

	b := bytes.NewBufferString("")
    cmd.SetOut(b)
	cmd.SetErr(b)

	err := cmd.Execute()
    assert.Error(t, err)

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

    assert.Contains(t, string(out), "Error:", "Wanted output to contain 'Error:'")
}

func newRootCmd() *cobra.Command {
	openapiCmd.AddCommand(validateOpenAPIV3Cmd)
	nyx.RootCmd.AddCommand(openapiCmd)

	return nyx.RootCmd 
}
