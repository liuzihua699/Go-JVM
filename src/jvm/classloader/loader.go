package classloader

import (
	"os"
	"strings"
)

const List_SEPARATOR = string(os.PathListSeparator)
const SEPARATOR = string(os.PathSeparator)
const CLASS_SEPARATOR = "/"

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
  create loader from path
  excute CreateCompositeLoader when path has recursion catalog. e.g com/zihua/User
  excute CreateWildcardLoader when path has wildcard. e.g com/* or *
  excute CreateZipLoader when path has .jar or .zip file
  else excute CreateDirLoader
*/
func CreateLoader(path string) ClassLoader {
	if strings.Contains(path, List_SEPARATOR) {
		return CreateCompositeLoader(path)
	}
	if strings.HasSuffix(path, "*") {
		return CreateWildcardLoader(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return CreateZipLoader(path)
	}
	return CreateDirLoader(path)
}
