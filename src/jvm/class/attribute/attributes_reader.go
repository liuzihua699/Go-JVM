package attribute

import "jvm/class/class_file_commons"

type AttributesStructReader struct {
	class_file_commons.Reader
}

func (fr *AttributesStructReader) ReadAttributeInfos() (uint16, *Attributes) {
	ret := new(Attributes)
	size := fr.ReadUint16()
	for i := 0; i < int(size); i++ {
		ret.Attributes = append(ret.Attributes, *new(AttributeInfo).ReadAttr(fr))
	}
	return size, ret
}
