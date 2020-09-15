package attribute

import "jvm/class/class_file_commons"

type AttributesStructReader struct {
	class_file_commons.Reader
}

func (fr *AttributesStructReader) ReadAttributeInfos() (uint16, *Attributes) {
	ret := new(Attributes)
	size := fr.ReadUint16()
	if size == 0 {
		ret = nil
	}
	for i := 0; i < int(size); i++ {
		ret.Attributes = append(ret.Attributes, ReadAttributeInfo(fr))
	}
	return size, ret
}
