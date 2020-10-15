package encrypt

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "encrypt"
	short = "Encrypt GPG messages like e.g. encrypted private keys."
	long  = `Encrypt GPG messages like e.g. encrypted private keys. Following conventions
and best practices should be respected, if not programmatically enforced.

    * Output files given by -o/--output must have the ".enc" suffix.

    * Passwords given by -p/--pass must at least be 64 characters long. If not
      given by command line flag an environment variable RED_GPG_PASS must be
      set in the process environment.

    * Encryption of specific file types like RSA deploy keys do not have to
      follow the convention of structured file system layout as described
      below.

Secure configuration management should follow a structured file system layout
as described below.

    sec
    ├── aws
    │   └── access
    │       ├── id.enc
    │       └── secret.enc
    └── docker
        ├── pass.enc
        └── user.enc

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
