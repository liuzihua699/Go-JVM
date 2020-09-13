package attribute

import "jvm/class/class_file_commons"

type Exceptions_attribute struct {
	AttributeInfo
	NumberOfExceptions  uint16
	ExceptionIndexTable []uint16 // size of NumberOfExceptions
}

func (e *Exceptions_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	e.NumberOfExceptions = reader.ReadUint16()
	size := e.NumberOfExceptions
	for i := 0; i < int(size); i++ {
		e.ExceptionIndexTable = append(e.ExceptionIndexTable, reader.ReadUint16())
	}
	return e
}
