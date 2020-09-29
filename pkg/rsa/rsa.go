package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/xh3b4sd/tracer"
	"golang.org/x/crypto/ssh"
)

// NewKeyPair generates a RSA public and private key. Both keys are PEM encoded.
// The first byte slice returned represents the private key. The second byte
// slice returned represents the public key.
func NewKeyPair() ([]byte, []byte, error) {
	pvk, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, tracer.Mask(err)
	}

	var pri []byte
	{
		var buf bytes.Buffer

		err := pem.Encode(
			&buf,
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(pvk),
			},
		)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}

		pri = buf.Bytes()
	}

	var pub []byte
	{
		key, err := ssh.NewPublicKey(&pvk.PublicKey)
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}

		pub = ssh.MarshalAuthorizedKey(key)
	}

	return pri, pub, nil
}
