package attribute

import "jvm/class/class_file_commons"

type ConstantValue_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	ConstantValueIndex uint16
}

func (c *ConstantValue_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	c.ConstantValueIndex = reader.ReadUint16()
	return c
}
