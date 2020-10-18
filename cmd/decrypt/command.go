package decrypt

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	name  = "decrypt"
	short = "Decrypt GPG messages like e.g. encrypted private keys."
	long  = `Decrypt GPG messages like e.g. encrypted private keys. Following conventions
and best practices should be respected, if not programmatically enforced.

    * Input files given by -i/--input must have the ".enc" suffix.

    * Output provided by -o/--output can either be a file on the file system or
      "-" to indicate to print to stdout.

    * Passwords given by -p/--pass can either be the password string itself or
      "-" to indicate to read from stdin. If not given by command line flag,
      an environment variable RED_GPG_PASS must be set in the process
      environment. Passwords must at least be 64 characters long.

    * Decryption of specific file types like RSA deploy keys do not have to
      follow the convention of structured file system layout as described
      below.

Secure configuration management should follow a structured file system layout
as described below. Usually a private repository should be created for the
sole purpose of secure configuration management. Secret rotation should be
implemented to rotate the password used for GPG operations. Secret rotations
should also be implemented to rotate the actually encrypted secrets.

    sec
    ├── aws
    │   └── access
    │       ├── id.enc
    │       └── secret.enc
    └── docker
        ├── pass.enc
        └── user.enc

The example below shows how to decrypt the secret data that is printed to
stdout.

    red decrypt -i key.enc -o - -p ********

The example below shows how to decrypt the GPG message read from a file on
the file system. The plain text secret is written to the configured output
file.

    red decrypt -i key.enc -o key.txt -p ********

The example below shows how to provide a password via stdin. Upon execution
of the command below the program will wait for any input made. Once the
[enter] key is pressed the command stops accepting input and uses the
provided password to decrypt the secret data.

    red decrypt -i key.enc -o key.txt -p -
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
