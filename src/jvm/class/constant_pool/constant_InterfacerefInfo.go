package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type ConstantInterfacerefInfo struct {
	TagInfo
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c *ConstantInterfacerefInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.ClassIndex = reader.ReadUint16()
	c.NameAndTypeIndex = reader.ReadUint16()
	return c
}
