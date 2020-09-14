package attribute

import (
	"jvm/class/class_file_commons"
)

type InnerClasses_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumberOfClasses    uint16
	Classes            []Classes
}

type Classes struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags uint16
}

func (c *Classes) readClasses(reader class_file_commons.Reader) Classes {
	c.InnerClassAccessFlags = reader.ReadUint16()
	c.OuterClassInfoIndex = reader.ReadUint16()
	c.InnerNameIndex = reader.ReadUint16()
	c.InnerClassAccessFlags = reader.ReadUint16()
	return *c
}

func (ia *InnerClasses_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	ia.NumberOfClasses = reader.ReadUint16()
	size := ia.NumberOfClasses
	for i := 0; i < int(size); i++ {
		ia.Classes = append(ia.Classes, new(Classes).readClasses(reader))
	}
	return ia
}
