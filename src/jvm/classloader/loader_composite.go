package classloader

type ComClassLoader struct {
}

func CreateCompositeLoader(path string) *ComClassLoader {
	return &ComClassLoader{}
}

func (c ComClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	panic("implement me")
}

func (c ComClassLoader) ToString() string {
	panic("implement me")
}
