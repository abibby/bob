package util

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/abibby/bob/slices"
)

func PackageRoot() (string, error) {
	cwd, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	current := cwd
	for {
		files, err := os.ReadDir(current)
		if err != nil {
			return "", err
		}

		_, ok := slices.Find(files, func(f fs.DirEntry) bool {
			return f.Name() == "go.mod"
		})
		if ok {
			return current, nil
		}

		if current == "/" {
			return cwd, nil
		}
		current = path.Dir(current)
	}
}
