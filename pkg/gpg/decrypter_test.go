package gpg

import (
	"testing"
)

func Test_Red_GPG_Decrypter_Decrypt(t *testing.T) {
	var err error

	var d *Decrypter
	{
		c := DecrypterConfig{
			Pass: "foo",
		}

		d, err = NewDecrypter(c)
		if err != nil {
			t.Fatal(err)
		}
	}

	expected := []byte("hello world")
	value := []byte(`-----BEGIN PGP SIGNATURE-----

wx4EBwMIbxESvPWYOgBgEmcsCe70T3fWYMXUAO/SBZHS4AHk6qQ8xikHnoCiBasb
8HKnT+FFOOCN4P3hujrg8+I+yH/84GfjXeQwavZMLtTgAeGQ8eCB4GXgguSqXR+p
l5T01NrWN5NQZM+H4ngibenh8GwA
=C3bS
-----END PGP SIGNATURE-----`) // "hello world"
	modified, err := d.Decrypt(value)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if string(modified) != string(expected) {
		t.Fatal("expected", expected, "got", modified)
	}
}

func Test_Red_GPG_Decrypter_Decrypt_Empty(t *testing.T) {
	var err error

	var d *Decrypter
	{
		c := DecrypterConfig{
			Pass: "foo",
		}

		d, err = NewDecrypter(c)
		if err != nil {
			t.Fatal(err)
		}
	}

	expected := []byte("")
	value := []byte("")
	modified, err := d.Decrypt(value)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if string(modified) != string(expected) {
		t.Fatal("expected", expected, "got", modified)
	}
}
