package classloader

import "path/filepath"

type DirClassLoader struct {
	absDir string
}

func CreateDirLoader(path string) *DirClassLoader {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirClassLoader{absDir}
}

func (d DirClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	panic("implement me")
}

func (d DirClassLoader) ToString() string {
	return "DirClassLoader"
}
