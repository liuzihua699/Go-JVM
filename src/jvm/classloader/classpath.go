package classloader

import (
	"errors"
	"jvm/commons"
	"strings"
)

type ClassPath struct {
	BootClassLoader *BaseLoader
	ExtClassLoader  *BaseLoader
	UserClassLoader *BaseLoader
}

func (c *ClassPath) InitClassPath(options commons.Cmd) error {
	err := errors.New("classpath initializer error!")

	lastDotIndex := strings.LastIndex(options.ClassName, ".")

	if lastDotIndex != -1 {
		c.BootClassLoader = CreateAppClassLoader(nil, options.BootClassPath+"\\*", "BootClassLoader")
		c.ExtClassLoader = CreateAppClassLoader(c.BootClassLoader, options.ExtClassPath+"\\*", "ExtClassLoader")
		c.UserClassLoader = CreateAppClassLoader(c.ExtClassLoader, options.ClassPath, "UserClassLoader")
	} else {
		c.BootClassLoader = CreateAppClassLoader(nil, options.BootClassPath, "BootClassLoader")
		c.ExtClassLoader = CreateAppClassLoader(c.BootClassLoader, options.ExtClassPath, "ExtClassLoader")
		c.UserClassLoader = CreateAppClassLoader(c.ExtClassLoader, options.ClassPath, "UserClassLoader")
	}

	if c.BootClassLoader.isNotReady() {
		err = errors.New("bootstrape classloader has not initializer.")
	} else if c.ExtClassLoader.isNotReady() {
		err = errors.New("extsions classloader has not initializer.")
	} else if c.UserClassLoader.isNotReady() {
		err = errors.New("user classpath classloader has not initializer.")
	} else {
		err = nil
	}
	return err
}

type BaseLoader struct {
	ParentClassLoader *BaseLoader
	currentLoader     ClassLoader
	name              string
}

func CreateAppClassLoader(parent *BaseLoader, path string, name string) *BaseLoader {
	return &BaseLoader{
		ParentClassLoader: parent,
		currentLoader:     CreateLoader(path),
		name:              name,
	}
}

func (l BaseLoader) hasParent() bool {
	return l.ParentClassLoader != nil
}

func (l BaseLoader) getParent() *BaseLoader {
	if l.hasParent() {
		return l.ParentClassLoader
	}
	return nil
}

func (l BaseLoader) getBaseLoader() *BaseLoader {
	return &l
}

func (l BaseLoader) isReady() bool {
	return l.currentLoader != nil
}

func (l BaseLoader) isNotReady() bool {
	return l.currentLoader == nil
}

func (l BaseLoader) setName(n string) {
	l.name = n
}
func (l BaseLoader) GetName() string {
	return l.name
}

// 用当前的类加载器加载
func (l BaseLoader) loadClass(className string) ([]byte, ClassLoader, error) {
	if l.currentLoader == nil {
		return nil, nil, errors.New("current ClassLoader not has initializer.")
	}
	return l.currentLoader.LoadClass(className)
}

// 使用委派机制的类加载器加载
func (l BaseLoader) ParentLoader(className string) ([]byte, ClassLoader, error, BaseLoader) {
	if l.hasParent() {
		data, loader, err, ll := l.ParentClassLoader.ParentLoader(className)
		if err == nil {
			return data, loader, err, ll
		}
		a, b, c := l.loadClass(className)
		return a, b, c, l
	}
	a, b, c := l.loadClass(className)
	return a, b, c, l
}
