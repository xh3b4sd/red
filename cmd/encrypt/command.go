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

    * Intput provided by -i/--input can either be a file on the file system or
      "-" to indicate to read from stdin.

    * Output files provided by -o/--output must have the ".enc" suffix.

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

The example below shows how to encrypt a secret that is provided via stdin.
Upon execution of the command below the program will wait for any input made.
Once the [enter] key is pressed the command stops accepting input and uses
the provided secret data to encrypt it. The encrypted GPG message is written
to the configured output file.

    red encrypt -i - -o key.enc -p ********

The example below shows how to encrypt the content of a file on the file
system. The encrypted GPG message is written to the configured output file.

    red encrypt -i key.txt -o key.enc -p ********
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
