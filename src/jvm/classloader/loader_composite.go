package classloader

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type ComClassLoader struct {
	paths   string
	listDir []string
	loaders []ClassLoader
}

/**
  This a composite loader with dir-loader and zip-loader
  e,g	re: (dir|zip)+;
*/
func CreateCompositeLoader(pathList string) *ComClassLoader {
	ret := new(ComClassLoader)
	ret.paths = pathList
	for _, path := range strings.Split(pathList, List_SEPARATOR) {
		absPath, fileAbsErr := filepath.Abs(path)
		ret.listDir = append(ret.listDir, absPath)
		if fileAbsErr != nil {
			continue
		}
		var err error
		if strings.HasSuffix(absPath, "*") {
			_, err = os.Stat(absPath[:len(absPath)-1])
		} else {
			_, err = os.Stat(absPath)
		}
		if err == nil {
			loader := CreateLoader(absPath)
			ret.loaders = append(ret.loaders, loader)
		}
	}
	return ret
}

func (c ComClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	var (
		data []byte = nil
		err         = errors.New("not fount file: " + className)
	)
	for _, loader := range c.loaders {
		data, _, err = loader.LoadClass(className)
		if err == nil {
			break
		}
	}
	return data, c, err
}

func (c ComClassLoader) ToString() string {
	return "ComClassLoader"
}
