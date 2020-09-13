package attribute

import (
	"jvm/class/class_file_commons"
)

type Code_attribute struct {
	AttributeInfo
	MaxStrack            uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionTable
	AttributesCount      uint16
	AttrInfo             Attributes
}

type ExceptionTable struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (c *Code_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	c.MaxStrack = reader.ReadUint16()
	c.MaxLocals = reader.ReadUint16()
	c.CodeLength = reader.ReadUint32()
	c.Code = reader.ReadBytes(c.CodeLength)
	c.ExceptionTableLength = reader.ReadUint16()
	size := c.ExceptionTableLength
	for i := 0; i < int(size); i++ {
		ex := new(ExceptionTable)
		ex.StartPc = reader.ReadUint16()
		ex.EndPc = reader.ReadUint16()
		ex.HandlerPc = reader.ReadUint16()
		ex.CatchType = reader.ReadUint16()
		c.ExceptionTable = append(c.ExceptionTable, *ex)
	}
	ar := AttributesStructReader{reader}
	a_count, a_infos := ar.ReadAttributeInfos()
	c.AttributesCount = a_count
	c.AttrInfo = *a_infos
	return c
}
