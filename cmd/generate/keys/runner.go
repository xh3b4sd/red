package keys

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/red/pkg/gpg"
	"github.com/xh3b4sd/red/pkg/pass"
	"github.com/xh3b4sd/red/pkg/rsa"
)

const (
	// Crt is the public key we generate.
	Crt = "id_rsa.pub"
	// Key is the private key we generate and encrypt using a password.
	Key = "id_rsa.enc"
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
	fmt.Println("Generating RSA keys and encryption password.")

	var err error

	var p string
	{
		p = pass.MustNew()
	}

	var e *gpg.Encrypter
	{
		c := gpg.EncrypterConfig{
			Pass: p,
		}

		e, err = gpg.NewEncrypter(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var crt []byte
	var key []byte
	{
		pri, pub, err := rsa.NewKeyPair()
		if err != nil {
			return tracer.Mask(err)
		}

		enc, err := e.Encrypt(pri)
		if err != nil {
			return tracer.Mask(err)
		}

		crt = pub
		key = enc
	}

	{
		p := r.flag.Directory

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cp string
	{
		cp = filepath.Join(r.flag.Directory, Crt)

		err = ioutil.WriteFile(cp, crt, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var kp string
	{
		kp = filepath.Join(r.flag.Directory, Key)

		err = ioutil.WriteFile(kp, key, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	fmt.Println()
	fmt.Printf("    password:       %s\n", p)
	fmt.Println()
	fmt.Printf("    private key:    %s\n", kp)
	fmt.Println()
	fmt.Printf("    public key:     %s\n", cp)
	fmt.Println()

	return nil
}
