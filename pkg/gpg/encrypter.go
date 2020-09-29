package gpg

import (
	"bytes"

	"github.com/xh3b4sd/tracer"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

type EncrypterConfig struct {
	Pass string
}

type Encrypter struct {
	pass string
}

func NewEncrypter(config EncrypterConfig) (*Encrypter, error) {
	if config.Pass == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Pass must not be empty", config)
	}

	e := &Encrypter{
		pass: config.Pass,
	}

	return e, nil
}

func (e *Encrypter) Encrypt(value []byte) ([]byte, error) {
	if len(value) == 0 {
		return nil, nil
	}

	buf := bytes.NewBuffer(nil)
	encoder, err := armor.Encode(buf, openpgp.SignatureType, nil)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	encrypter, err := openpgp.SymmetricallyEncrypt(encoder, []byte(e.pass), nil, nil)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	_, err = encrypter.Write(value)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	encrypter.Close()
	encoder.Close()

	return buf.Bytes(), nil
}
