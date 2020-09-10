package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/

type ConstantNameAndTypeInfo struct {
	TagInfo
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.NameIndex = reader.ReadUint16()
	c.DescriptorIndex = reader.ReadUint16()
	return c
}
