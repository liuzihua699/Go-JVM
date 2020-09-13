package attribute

import "jvm/class/class_file_commons"

/**
StackMapTable_attribute {
    u2              attribute_name_index;
    u4              attribute_length;
    u2              number_of_entries;
    stack_map_frame entries[number_of_entries];
}
*/
type StackMapTable_attribute struct {
	AttributeInfo
	NumberOfEntries uint16
	Entries         []StackMapFrame_Info
}

func (s *StackMapTable_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	s.NumberOfEntries = reader.ReadUint16()
	size := s.NumberOfEntries
	for i := 0; i < int(size); i++ {
		frameType := reader.ReadUint8()
		s.Entries = append(s.Entries, newStackMapFrame(frameType).ReadStackMapFrame(reader))
	}
	return s
}
