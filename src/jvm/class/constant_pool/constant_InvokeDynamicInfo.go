package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

type ConstantInvokeDynamicInfo struct {
	TagInfo
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (c *ConstantInvokeDynamicInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.BootstrapMethodAttrIndex = reader.ReadUint16()
	c.NameAndTypeIndex = reader.ReadUint16()
	return c
}
