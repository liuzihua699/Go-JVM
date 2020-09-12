package attribute

import "jvm/class/class_file_commons"

/**
e.g
	field_info {
		u2             access_flags;
		u2             name_index;
		u2             descriptor_index;
		u2             attributes_count;
		attribute_info attributes[attributes_count];
	}
*/
type MemberInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	AttributeInfos  []AttributeInfo
}

func (m *MemberInfo) ReadMem(reader class_file_commons.Reader) *MemberInfo {
	m.AccessFlags = reader.ReadUint16()
	m.NameIndex = reader.ReadUint16()
	m.DescriptorIndex = reader.ReadUint16()
	m.AttributesCount = reader.ReadUint16()
	size := int(m.AttributesCount)
	for i := 0; i < size; i++ {
		m.AttributeInfos = append(m.AttributeInfos, *new(AttributeInfo).ReadAttr(reader))
	}
	return m
}
