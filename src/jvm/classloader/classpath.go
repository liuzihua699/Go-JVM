package classloader

type ClassPath struct {
	BootClassLoader ClassLoader
	ExtClassLoader  ClassLoader
	UserClassLoader ClassLoader
}
