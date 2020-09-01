package classloader

type WildcardClassLoader struct {
}

func CreateWildcardLoader(path string) *WildcardClassLoader {
	return &WildcardClassLoader{}
}

func (w WildcardClassLoader) LoadClass(className string) ([]byte, ClassLoader, error) {
	panic("implement me")
}

func (w WildcardClassLoader) ToString() string {
	return "WildcardClassLoader"
}
