package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

type ConstantMethodHandleInfo struct {
	TagInfo
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (c *ConstantMethodHandleInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.ReferenceKind = reader.ReadUint8()
	c.ReferenceIndex = reader.ReadUint16()
	return c
}
