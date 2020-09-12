package fields

import (
	"jvm/class/attribute"
	"jvm/class/class_file_commons"
)

type FieldsStructReader struct {
	class_file_commons.Reader
}

func (fr *FieldsStructReader) ReadFieldsInfos() (int, *Fields) {
	ret := new(Fields)
	size := fr.ReadUint16()
	for i := 0; i < int(size); i++ {
		ret.MemberInfos = append(ret.MemberInfos, *new(attribute.MemberInfo).ReadMem(fr))
	}
	return int(size), ret
}
