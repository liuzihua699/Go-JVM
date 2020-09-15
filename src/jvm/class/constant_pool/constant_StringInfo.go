package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/

type ConstantStringInfo struct {
	TagInfo
	StringIndex uint16
}

func (c *ConstantStringInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.StringIndex = reader.ReadUint16()
	return c
}
