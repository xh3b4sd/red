package encrypt

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/red/pkg/env"
)

type flag struct {
	Input  string
	Output string
	Pass   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Input, "input", "i", "", "Input file to read the decrypted GPG message from.")
	cmd.Flags().StringVarP(&f.Output, "output", "o", "", "Output file to write the encrypted GPG message to.")
	cmd.Flags().StringVarP(&f.Pass, "pass", "p", "", "Password used for encryption of the GPG message.")
}

func (f *flag) Validate() error {
	{
		if f.Input == "" {
			return tracer.Maskf(invalidFlagError, "-i/--input must not be empty")
		}
	}

	{
		if f.Output == "" {
			return tracer.Maskf(invalidFlagError, "-o/--output must not be empty")
		}
		if !strings.HasSuffix(f.Output, ".enc") {
			return tracer.Maskf(invalidFlagError, "-o/--output must have suffix %#q", ".enc")
		}
	}

	{
		if f.Pass == "" {
			f.Pass = os.Getenv(env.RedGPGPass)
		}

		if f.Pass == "" {
			return tracer.Maskf(invalidFlagError, "-p/--pass must not be empty")
		}
		if len(f.Pass) < 64 {
			return tracer.Maskf(invalidFlagError, "-p/--pass must be at least 64 characters long")
		}
	}

	return nil
}
