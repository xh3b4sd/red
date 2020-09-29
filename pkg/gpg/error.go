package gpg

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidConfigError = &tracer.Error{
	Kind: "invalidConfigError",
}

func IsInvalidConfig(err error) bool {
	return errors.Is(err, invalidConfigError)
}

var decryptionFailedError = &tracer.Error{
	Kind: "decryptionFailedError",
}

func IsDecryptionFailed(err error) bool {
	return errors.Is(err, decryptionFailedError)
}
