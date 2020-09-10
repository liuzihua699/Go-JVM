package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

type ConstantMethodTypeInfo struct {
	TagInfo
	DescriptorIndex uint16
}

func (c *ConstantMethodTypeInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.DescriptorIndex = reader.ReadUint16()
	return c
}
