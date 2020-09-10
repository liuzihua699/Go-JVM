package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantLongInfo struct {
	TagInfo
	HighBytes uint32
	LowBytes  uint32
}

func (c *ConstantLongInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.HighBytes = reader.ReadUint32()
	c.LowBytes = reader.ReadUint32()
	return c
}
