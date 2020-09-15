package constant_pool

import "jvm/class/class_file_commons"

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

type ConstantClassInfo struct {
	TagInfo
	NameIndex uint16
}

func (c *ConstantClassInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.NameIndex = reader.ReadUint16()
	return c
}
