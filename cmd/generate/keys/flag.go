package keys

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Directory string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Directory, "directory", "d", ".", "Directory to write key files to.")
}

func (f *flag) Validate() error {
	if f.Directory == "" {
		return tracer.Maskf(invalidFlagError, "-d/--directory must not be empty")
	}

	return nil
}
