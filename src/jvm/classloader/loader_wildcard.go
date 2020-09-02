package classloader

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

/**
  order：
    1. dirListPath
	2. jarListPath
	3. zipListPath
*/
type WildcardClassLoader struct {
	basePath    string
	dirListPath []string
	jarListPath []string
	zipListPath []string
	loaders     []ClassLoader
}

func CreateWildcardLoader(path string) *WildcardClassLoader {
	ret := new(WildcardClassLoader)
	basePath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	basePath = basePath[:len(basePath)-1]
	ret.basePath = basePath

	_, err = os.Stat(basePath)
	if err != nil || os.IsNotExist(err) {
		panic(err)
	}

	walkFun := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			ret.dirListPath = append(ret.dirListPath, path)
		} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			ret.jarListPath = append(ret.jarListPath, path)
		} else if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			ret.zipListPath = append(ret.zipListPath, path)
		}
		return err
	}

	err = filepath.Walk(basePath, walkFun)
	if err != nil {
		panic(err)
	}

	return ret
}

func (w WildcardClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	lastDotIndex := strings.LastIndex(className, ".")
	relativeClassPath := strings.ReplaceAll(className[:lastDotIndex], ".", SEPARATOR)
	for _, path := range w.dirListPath {
		if strings.Contains(path, relativeClassPath) {
			w.loaders = append(w.loaders, CreateLoader(path))
		}
	}

	// com.zihua.HelloWorld
	// [:last] = com.zihua
	// [last+1:] = HelloWorld
	var (
		data []byte = nil
		err  error  = errors.New("Not fount class exception!")
	)
	for _, loader := range w.loaders {
		data, _, err = loader.LoadClass(className[lastDotIndex+1:])
		if err == nil && data != nil {
			return data, w, err
		}
	}

	for _, path := range w.jarListPath {
		data, _, err = CreateLoader(path).LoadClass(className)
		if err == nil && data != nil {
			return data, w, err
		}
	}

	for _, path := range w.zipListPath {
		data, _, err = CreateLoader(path).LoadClass(className)
		if err == nil && data != nil {
			return data, w, err
		}
	}

	return nil, w, err
}

func (w WildcardClassLoader) ToString() string {
	return "WildcardClassLoader"
}
