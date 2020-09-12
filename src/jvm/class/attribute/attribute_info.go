package attribute

import "jvm/class/class_file_commons"

/**
e.g
	attribute_info {
		u2 attribute_name_index;
		u4 attribute_length;
		u1 info[attribute_length];
	}
*/
/**
根据官方说明，info的类型应为attribute_name_index索引的类型，一共23种
*/
type AttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	AttrInfo           []AttrInfo
}

func (a *AttributeInfo) ReadAttr(reader class_file_commons.Reader) *AttributeInfo {
	a.AttributeNameIndex = reader.ReadUint16()
	a.AttributeLength = reader.ReadUint32()
	a.AttrInfo = append(a.AttrInfo, newAttributeInfo(*a).ReadAttrInfo(reader))
	return a
}
