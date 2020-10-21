# red

Managing gpg messages and rsa keys.



### Decryption

```
$ red decrypt -h
Decrypt GPG messages like e.g. encrypted private keys. Following conventions
and best practices should be respected, if not programmatically enforced.

    * Input provided by -i/--input can either be a file or directory on the
      file system. Files must have the ".enc" suffix. Directories will be
      traversed to find all files having the ".enc" suffix. All found files
      will be decrypted and serialized into a structured JSON object
      according to the file system structure.

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

    .
    ├── aws
    │   ├── accessid.enc
    │   └── secretid.enc
    └── docker
        ├── password.enc
        ├── registry.enc
        └── username.enc

When decrypting all secrets within a directory the serialized JSON object
according to the example file system structure shown above will look like the
following.

    {
        "aws.accessid": "...",
        "aws.secretid": "...",
        "docker.password": "...",
        "docker.registry": "...",
        "docker.username": "..."
    }

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

The example below shows how to decrypt all secrets within a directory and
write the serialized JSON object to stdout. Note that the command below
defines the "-s" flag, which makes it operate in silent mode for scripting
and suppresses labels.

    red decrypt -i . -o - -p ******** -s

Usage:
  red decrypt [flags]

Flags:
  -h, --help            help for decrypt
  -i, --input string    Input file to read the encrypted GPG message from.
  -o, --output string   Output file to write the decrypted GPG message to.
  -p, --pass string     Password used for decryption of the GPG message.
  -s, --silent          Silent mode for scripting to suppress labels.
```



### Encryption

```
$ red encrypt -h
Encrypt GPG messages like e.g. encrypted private keys. Following conventions
and best practices should be respected, if not programmatically enforced.

    * Intput provided by -i/--input can either be a file on the file system or
      "-" to indicate to read from stdin.

    * Output files provided by -o/--output must have the ".enc" suffix.

    * Passwords given by -p/--pass can either be the password string itself or
      "-" to indicate to read from stdin. If not given by command line flag,
      an environment variable RED_GPG_PASS must be set in the process
      environment. Passwords must at least be 64 characters long.

    * Encryption of specific file types like RSA deploy keys do not have to
      follow the convention of structured file system layout as described
      below.

Secure configuration management should follow a structured file system layout
as described below. Usually a private repository should be created for the
sole purpose of secure configuration management. Secret rotation should be
implemented to rotate the password used for GPG operations. Secret rotations
should also be implemented to rotate the actually encrypted secrets.

    .
    ├── aws
    │   ├── accessid.enc
    │   └── secretid.enc
    └── docker
        ├── password.enc
        ├── registry.enc
        └── username.enc

The example below shows how to encrypt the secret data that is provided via
stdin. Upon execution of the command below the program will wait for any
input made. Once the [enter] key is pressed the command stops accepting input
and uses the provided secret data to encrypt it. The encrypted GPG message is
written to the configured output file.

    red encrypt -i - -o key.enc -p ********

The example below shows how to encrypt the content of a file on the file
system. The encrypted GPG message is written to the configured output file.

    red encrypt -i key.txt -o key.enc -p ********

The example below shows how to provide a password via stdin. Upon execution
of the command below the program will wait for any input made. Once the
[enter] key is pressed the command stops accepting input and uses the
provided password to encrypt the secret data.

    red encrypt -i key.txt -o key.enc -p -

Usage:
  red encrypt [flags]

Flags:
  -h, --help            help for encrypt
  -i, --input string    Input file to read the decrypted GPG message from.
  -o, --output string   Output file to write the encrypted GPG message to.
  -p, --pass string     Password used for encryption of the GPG message.
```
