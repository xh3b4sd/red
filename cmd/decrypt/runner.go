package decrypt

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/gpg"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/red/pkg/file"
)

type runner struct {
	flag   *flag
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Stdin()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.flag.Validate()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) decFromDir(d *gpg.Decrypter) ([]byte, error) {
	// The first thing we need to do is to lookup the absolut paths of all
	// encrypted files. The resulting list of files might look like below.
	//
	//     /Users/xh3b4sd/projects/xh3b4sd/sec/docker/pass.enc
	//     /Users/xh3b4sd/projects/xh3b4sd/sec/docker/regi.enc
	//     /Users/xh3b4sd/projects/xh3b4sd/sec/docker/user.enc
	//
	var files []string
	{
		walkFunc := func(p string, i os.FileInfo, err error) error {
			if err != nil {
				return tracer.Mask(err)
			}

			if i.IsDir() && i.Name() == ".git" {
				return filepath.SkipDir
			}

			if i.IsDir() && i.Name() == ".github" {
				return filepath.SkipDir
			}

			// We do not want to track directories. We are interested in
			// directories containing specific files.
			if i.IsDir() {
				return nil
			}

			// We do not want to track files with the wrong extension. We are
			// interested in encrypted files having the ".enc" extension.
			if filepath.Ext(i.Name()) != ".enc" {
				return nil
			}

			files = append(files, p)

			return nil
		}

		err := afero.Walk(afero.NewOsFs(), r.flag.Input, walkFunc)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	m := map[string]string{}
	for _, f := range files {
		enc, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		dec, err := d.Decrypt(enc)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		f = strings.TrimPrefix(f, mustAbs(r.flag.Input)+"/")
		f = strings.TrimSuffix(f, filepath.Ext(f))
		f = strings.ReplaceAll(f, "/", ".")

		m[f] = string(dec)
	}

	var dec []byte
	{
		b, err := json.Marshal(m)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		dec = []byte(fmt.Sprintf("%s\n", b))
	}

	return dec, nil
}

func (r *runner) decFromFile(d *gpg.Decrypter) ([]byte, error) {
	var enc []byte
	{
		p := r.flag.Input

		b, err := ioutil.ReadFile(p)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		enc = b
	}

	var dec []byte
	{
		b, err := d.Decrypt(enc)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		// For convenience we want to append a new line at the end of the
		// decrypted secret. This helps printing plain text secrets to stdout as
		// well as writing them to files on the file system.
		dec = []byte(fmt.Sprintf("%s\n", b))
	}

	return dec, nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var d *gpg.Decrypter
	{
		c := gpg.DecrypterConfig{
			Pass: r.flag.Pass,
		}

		d, err = gpg.NewDecrypter(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var dec []byte
	if file.IsDir(r.flag.Input) {
		dec, err = r.decFromDir(d)
		if err != nil {
			return tracer.Mask(err)
		}
	} else {
		dec, err = r.decFromFile(d)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if r.flag.Output == "-" {
		if !r.flag.Silent {
			fmt.Print("-o/--output: ")
		}
		fmt.Printf("%s", dec)
	} else {
		p := r.flag.Output

		err = os.MkdirAll(filepath.Dir(p), os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}

		err = ioutil.WriteFile(p, dec, 0600)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func mustAbs(p string) string {
	abs, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}

	return abs
}
