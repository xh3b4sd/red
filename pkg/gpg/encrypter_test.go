package gpg

import (
	"strings"
	"testing"
)

func Test_Red_GPG_Encrypter_Encrypt(t *testing.T) {
	var err error

	var e *Encrypter
	{
		c := EncrypterConfig{
			Pass: "foo",
		}

		e, err = NewEncrypter(c)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Note that we cannot predict the outcome of the encryption due to the fact
	// how GPG works. Below we can only assume to have a proper GPG message by
	// comparing its length and verify the prefix and suffix of the GPG message is
	// as expected.
	expected := []byte(`-----BEGIN PGP SIGNATURE-----

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxx
-----END PGP SIGNATURE-----`) // "hello world"
	value := []byte("hello world")
	modified, err := e.Encrypt(value)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	le := len(string(expected))
	lm := len(string(modified))
	if lm != le {
		t.Fatal("expected", le, "got", lm)
	}
	if !strings.HasPrefix(string(modified), "-----BEGIN PGP SIGNATURE-----\n\n") {
		t.Fatal("expected", true, "got", false)
	}
	if !strings.HasSuffix(string(modified), "\n-----END PGP SIGNATURE-----") {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Red_GPG_Encrypter_Encrypt_Empty(t *testing.T) {
	var err error

	var e *Encrypter
	{
		c := EncrypterConfig{
			Pass: "foo",
		}

		e, err = NewEncrypter(c)
		if err != nil {
			t.Fatal(err)
		}
	}

	expected := []byte("")
	value := []byte("")
	modified, err := e.Encrypt(value)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if string(modified) != string(expected) {
		t.Fatal("expected", expected, "got", modified)
	}
}
