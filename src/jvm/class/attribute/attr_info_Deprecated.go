package attribute

import "jvm/class/class_file_commons"

type Deprecated_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
}

func (d Deprecated_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	return d
}
