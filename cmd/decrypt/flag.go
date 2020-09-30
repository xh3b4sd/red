package decrypt

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Input  string
	Output string
	Pass   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Input, "input", "i", "id_rsa.enc", "Input file to read the encrypted GPG message from.")
	cmd.Flags().StringVarP(&f.Output, "output", "o", "id_rsa", "Output file to write the decrypted GPG message to.")
	cmd.Flags().StringVarP(&f.Pass, "pass", "p", "********", "Password used for decryption of the GPG message.")
}

func (f *flag) Validate() error {
	if f.Input == "" {
		return tracer.Maskf(invalidFlagError, "-i/--input must not be empty")
	}
	if f.Output == "" {
		return tracer.Maskf(invalidFlagError, "-o/--output must not be empty")
	}
	if f.Pass == "" {
		return tracer.Maskf(invalidFlagError, "-p/--pass must not be empty")
	}

	return nil
}
