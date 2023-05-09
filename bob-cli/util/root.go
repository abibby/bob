package util

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/abibby/bob/slices"
	"golang.org/x/mod/modfile"
)

var ErrNoPackage = fmt.Errorf("not in a go package")

func PackageRoot(from string) (string, error) {
	cwd, err := filepath.Abs(from)
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
			return "", ErrNoPackage
		}
		current = path.Dir(current)
	}
}

type PackageInfo struct {
	PackageRoot string
	ImportPath  string
}

func PkgInfo(from string) (*PackageInfo, error) {
	root, err := PackageRoot(from)
	if err != nil {
		return nil, err
	}

	modFile := path.Join(root, "go.mod")
	b, err := os.ReadFile(modFile)
	if err != nil {
		return nil, err
	}

	m, err := modfile.ParseLax(modFile, b, nil)
	if err != nil {
		return nil, err
	}

	return &PackageInfo{
		PackageRoot: root,
		ImportPath:  m.Module.Syntax.Token[1],
	}, nil
}
