package attribute

import "jvm/class/class_file_commons"

type LineNumberTable_attribute struct {
	AttributeNameIndex    uint16
	AttributeLength       uint32
	LineNumberTableLength uint16
	LineNumberTable       []LineNumberTable1
}

type LineNumberTable1 struct {
	StartPc    uint16
	LineNumber uint16
}

func (l *LineNumberTable_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	l.LineNumberTableLength = reader.ReadUint16()
	size := l.LineNumberTableLength
	for i := 0; i < int(size); i++ {
		t := new(LineNumberTable1)
		t.StartPc = reader.ReadUint16()
		t.LineNumber = reader.ReadUint16()
		l.LineNumberTable = append(l.LineNumberTable, *t)
	}
	return l
}
