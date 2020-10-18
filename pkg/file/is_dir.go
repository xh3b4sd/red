package file

import "os"

func IsDir(file string) bool {
	if !Exists(file) {
		return false
	}

	i, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	return i.IsDir()
}
