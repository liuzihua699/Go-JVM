package attribute

import (
	"jvm/class/class_file_commons"
)

/**
The attribute reader not found format in implemented attribute-table, used it.
*/
type Default_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []byte
}

func (b *Default_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	b.Info = reader.ReadBytes(b.AttributeLength)
	return b
}
