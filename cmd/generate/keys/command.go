package keys

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "keys"
	short = "Generate RSA keys for e.g. github workflows."
	long  = `Generate RSA keys for e.g. github workflows. The generated private key is GPG
encrypted. It can be decrypted using the generated GPG password. The
generated public key is provided in plain text. See example usage and output
below.

    $ red generate keys -d .github/asset
    Generating RSA keys and encryption password.

        password:       *_E(O_r5x5/:aqy&l,QY0:sGPB^Vupd.(oeJA@xw{1,av$]J,@Bc&sjcr)jsB{s2

        private key:    .github/asset/id_rsa.enc

        public key:     .github/asset/id_rsa.pub

`
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var c *cobra.Command
	{
		f := &flag{}

		r := &runner{
			flag:   f,
			logger: config.Logger,
		}

		c = &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
			RunE:  r.Run,
		}

		f.Init(c)
	}

	return c, nil
}
