package project

var (
	description = "Managing gpg messages and rsa keys."
	gitSHA      = "n/a"
	name        = "pag"
	source      = "https://github.com/xh3b4sd/red"
	version     = "n/a"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
