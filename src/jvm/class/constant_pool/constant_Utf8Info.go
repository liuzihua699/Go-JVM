package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/

type ConstantUtf8Info struct {
	TagInfo
	Length uint16
	Bytes  []byte
}

func (c *ConstantUtf8Info) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.Length = reader.ReadUint16()
	c.Bytes = reader.ReadBytes(uint32(c.Length))
	return c
}
