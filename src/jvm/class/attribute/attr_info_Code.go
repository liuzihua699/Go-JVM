package attribute

import (
	"jvm/class/class_file_commons"
)

/**
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type Code_attribute struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	MaxStrack            uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionTable
	AttributesCount      uint16
	Attributes           Attributes
}

type ExceptionTable struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (e *ExceptionTable) readExceptionTable(reader class_file_commons.Reader) ExceptionTable {
	e.StartPc = reader.ReadUint16()
	e.EndPc = reader.ReadUint16()
	e.HandlerPc = reader.ReadUint16()
	e.CatchType = reader.ReadUint16()
	return *e
}

func (c *Code_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	c.MaxStrack = reader.ReadUint16()
	c.MaxLocals = reader.ReadUint16()
	c.CodeLength = reader.ReadUint32()
	c.Code = reader.ReadBytes(c.CodeLength)
	c.ExceptionTableLength = reader.ReadUint16()
	size := c.ExceptionTableLength
	for i := 0; i < int(size); i++ {
		c.ExceptionTable = append(c.ExceptionTable, new(ExceptionTable).readExceptionTable(reader))
	}
	ar := AttributesStructReader{reader}
	a_count, a_infos := ar.ReadAttributeInfos()
	c.AttributesCount = a_count
	c.Attributes = *a_infos
	return c
}
