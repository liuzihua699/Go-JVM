package attribute

import "jvm/class/class_file_commons"

type Signature_attribute struct {
	AttributeInfo
	Signature_index uint16
}

func (s Signature_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	s.Signature_index = reader.ReadUint16()
	return s
}
