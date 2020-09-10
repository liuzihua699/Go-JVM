package class_file_commons

import (
	"encoding/binary"
)

type Reader interface {
	ReadUint8() uint8            // 读取1字节，对应java中的bit
	ReadUint16() uint16          // 读取2字节，对应java中的short
	ReadUint32() uint32          // 读取4字节，对应java中的int
	ReadUint64() uint64          // 读取8字节，对应java中的long
	ReadBytes(len uint32) []byte // 读取定长字节数
	Clear()
}

type ClassFileReader struct {
	ByteStream *[]byte // 字节流地址
}

func (c *ClassFileReader) ReadUint8() uint8 {
	ret := (*c.ByteStream)[0]
	*c.ByteStream = (*c.ByteStream)[1:]
	return ret
}

func (c *ClassFileReader) ReadUint16() uint16 {
	ret := binary.BigEndian.Uint16(*c.ByteStream)
	*c.ByteStream = (*c.ByteStream)[2:]
	return ret
}

func (c *ClassFileReader) ReadUint32() uint32 {
	ret := binary.BigEndian.Uint32(*c.ByteStream)
	*c.ByteStream = (*c.ByteStream)[4:]
	return ret
}

func (c *ClassFileReader) ReadUint64() uint64 {
	ret := binary.BigEndian.Uint64(*c.ByteStream)
	*c.ByteStream = (*c.ByteStream)[8:]
	return ret
}

func (c *ClassFileReader) ReadBytes(length uint32) []byte {
	ret := (*c.ByteStream)[:length]
	*c.ByteStream = (*c.ByteStream)[length:]
	return ret
}

func (c *ClassFileReader) Clear() {
	c.ByteStream = nil
}

//func (c *ClassFileReader) ReadClassFile() (*ClassFile, error) {
//	return nil,nil
//}
