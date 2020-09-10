package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/

type ConstantIntegerInfo struct {
	TagInfo
	Bytes uint32
}

func (c *ConstantIntegerInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.Bytes = reader.ReadUint32()
	return c
}
