package encrypt

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/gpg"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type runner struct {
	flag   *flag
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Stdin()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.flag.Validate()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var e *gpg.Encrypter
	{
		c := gpg.EncrypterConfig{
			Pass: r.flag.Pass,
		}

		e, err = gpg.NewEncrypter(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var dec []byte
	if r.flag.inputFromStdin {
		dec = []byte(r.flag.Input)
	} else {
		p := r.flag.Input

		b, err := ioutil.ReadFile(p)
		if err != nil {
			return tracer.Mask(err)
		}

		dec = b
	}

	var enc []byte
	{
		b, err := e.Encrypt(dec)
		if err != nil {
			return tracer.Mask(err)
		}

		// For convenience we want to append a new line at the end of the
		// encrypted secret. This helps writing GPG messages to files on the
		// file system upon human inspection.
		enc = []byte(fmt.Sprintf("%s\n", b))
	}

	{
		p := r.flag.Output

		err = os.MkdirAll(filepath.Dir(p), os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}

		err = ioutil.WriteFile(p, enc, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
