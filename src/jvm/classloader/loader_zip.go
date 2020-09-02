package classloader

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ZipClassLoader struct {
	absZipPath string
}

func CreateZipLoader(path string) *ZipClassLoader {
	absZipPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(absZipPath)
	if err != nil || os.IsNotExist(err) {
		panic(err)
	}
	if !(strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") ||
		strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR")) {
		panic("Not support this type.")
	}

	return &ZipClassLoader{
		absZipPath: absZipPath,
	}
}

func (z ZipClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	zipFile, err := zip.OpenReader(z.absZipPath)
	defer zipFile.Close()
	if err != nil {
		return nil, z, err
	}

	className = strings.ReplaceAll(className, ".", CLASS_SEPARATOR) + ".class"

	var classFile *zip.File
	for _, file := range zipFile.File {
		if file.Name == className {
			classFile = file
			break
		}
	}

	if classFile == nil {
		return nil, z, errors.New("Not fount class exception!")
	}

	rc, err := classFile.Open()
	defer rc.Close()

	if err != nil {
		return nil, z, err
	}

	data, err := ioutil.ReadAll(rc)
	return data, z, err
}

func (z ZipClassLoader) ToString() string {
	return "ZipClassLoader"
}
