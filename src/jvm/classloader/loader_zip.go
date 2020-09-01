package classloader

type ZipClassLoader struct {
}

func CreateZipLoader(path string) *ZipClassLoader {
	return &ZipClassLoader{}
}

func (z ZipClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	panic("implement me")
}

func (z ZipClassLoader) ToString() string {
	panic("implement me")
}
