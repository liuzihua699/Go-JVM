package classloader

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type DirClassLoader struct {
	absDir string
}

func CreateDirLoader(path string) *DirClassLoader {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(absDir)
	if err != nil || os.IsNotExist(err) {
		panic(err)
	}

	return &DirClassLoader{absDir}
}

func (d DirClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	if strings.HasPrefix(className, ".class") {
		return nil, d, errors.New("Not fount this file.")
	}
	className = strings.ReplaceAll(className, ".", SEPARATOR) + ".class"
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

func (d DirClassLoader) ToString() string {
	return "DirClassLoader"
}
