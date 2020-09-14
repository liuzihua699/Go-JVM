package attribute

import "jvm/class/class_file_commons"

type LineNumberTable_attribute struct {
	AttributeNameIndex    uint16
	AttributeLength       uint32
	LineNumberTableLength uint16
	LineNumberTables      []LineNumberTable
}

type LineNumberTable struct {
	StartPc    uint16
	LineNumber uint16
}

func (l *LineNumberTable) readLineNumberTable(reader class_file_commons.Reader) LineNumberTable {
	l.StartPc = reader.ReadUint16()
	l.LineNumber = reader.ReadUint16()
	return *l
}

func (l *LineNumberTable_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	l.LineNumberTableLength = reader.ReadUint16()
	size := l.LineNumberTableLength
	for i := 0; i < int(size); i++ {
		l.LineNumberTables = append(l.LineNumberTables, new(LineNumberTable).readLineNumberTable(reader))
	}
	return l
}
