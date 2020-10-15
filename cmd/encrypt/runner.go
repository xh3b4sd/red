package encrypt

import (
	"context"
	"io/ioutil"

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

	err := r.flag.Validate()
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

	var enc []byte
	{
		p := r.flag.Input

		b, err := ioutil.ReadFile(p)
		if err != nil {
			return tracer.Mask(err)
		}

		enc = b
	}

	var dec []byte
	{
		b, err := e.Encrypt(enc)
		if err != nil {
			return tracer.Mask(err)
		}

		dec = b
	}

	{
		p := r.flag.Output

		err = ioutil.WriteFile(p, dec, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
