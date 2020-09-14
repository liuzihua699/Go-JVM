package attribute

import "jvm/class/class_file_commons"

type Signature_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Signature_index    uint16
}

func (s Signature_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	s.Signature_index = reader.ReadUint16()
	return s
}
