package attribute

import (
	"jvm/class/class_file_commons"
	"jvm/class/constant_pool"
)

/**
used for Fields、Methods、Interfaces and more than.
*/
type MemberInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttrInfo
}

func (m *MemberInfo) ReadMem(reader class_file_commons.Reader) *MemberInfo {
	m.AccessFlags = reader.ReadUint16()
	m.NameIndex = reader.ReadUint16()
	m.DescriptorIndex = reader.ReadUint16()
	m.AttributesCount = reader.ReadUint16()
	size := int(m.AttributesCount)
	for i := 0; i < size; i++ {
		m.Attributes = append(m.Attributes, ReadAttributeInfo(reader))
	}
	return m
}

func (m MemberInfo) GetName() string {
	return constant_pool.GetFileConstantPool().Utf8Map[m.NameIndex]
}

func (m MemberInfo) GetDescripor() string {
	return constant_pool.GetFileConstantPool().Utf8Map[m.DescriptorIndex]
}
