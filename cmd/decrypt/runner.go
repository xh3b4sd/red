package decrypt

import (
	"context"
	"fmt"
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

	var d *gpg.Decrypter
	{
		c := gpg.DecrypterConfig{
			Pass: r.flag.Pass,
		}

		d, err = gpg.NewDecrypter(c)
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
		b, err := d.Decrypt(enc)
		if err != nil {
			return tracer.Mask(err)
		}

		// For convenience we want to append a new line at the end of the
		// decrypted secret. This helps printing plain text secrets to stdout as
		// well as writing them to files on the file system.
		dec = []byte(fmt.Sprintf("%s\n", b))
	}

	if r.flag.Output == "-" {
		fmt.Printf("%s", dec)
	} else {
		p := r.flag.Output

		err = ioutil.WriteFile(p, dec, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
