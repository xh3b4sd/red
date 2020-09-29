package gpg

import (
	"bytes"
	"io/ioutil"

	"github.com/xh3b4sd/tracer"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

type DecrypterConfig struct {
	Pass string
}

type Decrypter struct {
	pass string
}

func NewDecrypter(config DecrypterConfig) (*Decrypter, error) {
	if config.Pass == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Pass must not be empty", config)
	}

	d := &Decrypter{
		pass: config.Pass,
	}

	return d, nil
}

func (d *Decrypter) Decrypt(value []byte) ([]byte, error) {
	if len(value) == 0 {
		return nil, nil
	}

	buf := bytes.NewBuffer(value)
	decoder, err := armor.Decode(buf)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	promptFunc := func() func([]openpgp.Key, bool) ([]byte, error) {
		retried := false
		return func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
			if !retried {
				retried = true
				return []byte(d.pass), nil
			}
			return nil, tracer.Mask(decryptionFailedError)
		}
	}()
	details, err := openpgp.ReadMessage(decoder.Body, nil, promptFunc, nil)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	b, err := ioutil.ReadAll(details.UnverifiedBody)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}
