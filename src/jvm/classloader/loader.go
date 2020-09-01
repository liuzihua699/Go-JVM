package classloader

import (
	"os"
	"strings"
)

const SEPARATOR = string(os.PathListSeparator)

/**
  From 《Java Virtual Machine Specification》, I implement 4 type class-loader,
  They generate the corresponding class loader based on the path:
	1. DirClassLoader		e.g com/zihua/user/
	2. ZipClassLoader		e.g %JAVA_HOME%/jre/lib/rt.jar
	3. WildcardClassLoader	e.g com/zihua/*
	4. ComClassLoader		e.g com/zihua/system/;com/zihua/user/
*/
type ClassLoader interface {
	LoadClass(className string) ([]byte, ClassLoader, error)
	ToString() string
}

/**
  create entry from path
  excute createCompositeLoader when path has recursion catalog. e.g com/zihua/User
  excute createWildcardLoader when path has wildcard. e.g com/* or *
  excute createZipLoader when path has .jar or .zip file
  else excute createDirLoader
*/
func CreateLoader(path string) ClassLoader {
	if strings.Contains(path, SEPARATOR) {
		return CreateCompositeLoader(path)
	}
	if strings.HasSuffix(path, "*") {
		return CreateWildcardLoader(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		return CreateZipLoader(path)
	}
	return CreateDirLoader(path)
}
