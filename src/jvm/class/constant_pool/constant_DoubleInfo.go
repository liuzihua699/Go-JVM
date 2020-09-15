package constant_pool

import (
	"jvm/class/class_file_commons"
)

/**
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantDoubleInfo struct {
	TagInfo
	HighBytes uint32
	LowBytes  uint32
}

func (c *ConstantDoubleInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.HighBytes = reader.ReadUint32()
	c.LowBytes = reader.ReadUint32()
	return c
}
