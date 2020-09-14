package attribute

import (
	"jvm/class/class_file_commons"
)

type SourceFile_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	SourceFileIndex    uint16
}

func (s *SourceFile_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	s.SourceFileIndex = reader.ReadUint16()
	return s
}
