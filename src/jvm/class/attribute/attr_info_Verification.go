package attribute

import (
	"jvm/class/class_file_commons"
)

/**
union verification_type_info {
    Top_variable_info; // 0
    Integer_variable_info; // 1
    Float_variable_info; // 2
    Long_variable_info; // 4
    Double_variable_info; // 3
    Null_variable_info; // 5
    UninitializedThis_variable_info; // 6
    Object_variable_info; // 7
    Uninitialized_variable_info; // 8
}
*/

const (
	ITEM_Top               = 0
	ITEM_Integer           = 1
	ITEM_Float             = 2
	ITEM_Double            = 3
	ITEM_Long              = 4
	ITEM_Null              = 5
	ITEM_UninitializedThis = 6
	ITEM_Object            = 7
	ITEM_Uninitialized     = 8
)

type VerificationTypeInfo interface {
	ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo
}

func newVerificationTypeInfo(tag uint8) VerificationTypeInfo {
	switch tag {
	case ITEM_Top:
		return &TopVariableInfo{Tag: tag}
	case ITEM_Integer:
		return &IntegerVariableInfo{Tag: tag}
	case ITEM_Float:
		return &FloatVariableInfo{Tag: tag}
	case ITEM_Null:
		return &NullVariableInfo{Tag: tag}
	case ITEM_UninitializedThis:
		return &UninitializedThisVariableInfo{Tag: tag}
	case ITEM_Object:
		return &ObjectVariableInfo{Tag: tag}
	case ITEM_Uninitialized:
		return &UninitializedVariableInfo{Tag: tag}
	case ITEM_Long:
		return &LongVariableInfo{Tag: tag}
	case ITEM_Double:
		return &DoubleVariableInfo{Tag: tag}
	}

	return nil
}

/**
Top_variable_info {
    u1 tag = ITEM_Top; // 0
}
*/
type TopVariableInfo struct {
	Tag uint8
}

func (t TopVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return t
}

/**
Integer_variable_info {
    u1 tag = ITEM_Integer; // 1
}
*/
type IntegerVariableInfo struct {
	Tag uint8
}

func (i IntegerVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return i
}

/**
Float_variable_info {
    u1 tag = ITEM_Float; // 2
}
*/
type FloatVariableInfo struct {
	Tag uint8
}

func (f FloatVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return f
}

/**
Null_variable_info {
    u1 tag = ITEM_Null; // 5
}
*/
type NullVariableInfo struct {
	Tag uint8
}

func (n NullVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return n
}

/**
UninitializedThis_variable_info {
    u1 tag = ITEM_UninitializedThis; // 6
}
*/
type UninitializedThisVariableInfo struct {
	Tag uint8
}

func (u UninitializedThisVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return u
}

/**
Object_variable_info {
    u1 tag = ITEM_Object; // 7
	u2 cpool_index;
}
*/
type ObjectVariableInfo struct {
	Tag         uint8
	Cpool_index uint16
}

func (o *ObjectVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	o.Cpool_index = reader.ReadUint16()
	return o
}

/**
Uninitialized_variable_info {
    u1 tag = ITEM_Uninitialized; // 8
	u2 offset;
}
*/
type UninitializedVariableInfo struct {
	Tag    uint8
	offset uint16
}

func (u *UninitializedVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	u.offset = reader.ReadUint16()
	return u
}

/**
Long_variable_info {
    u1 tag = ITEM_Long; // 4
}
*/
type LongVariableInfo struct {
	Tag uint8
}

func (l LongVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return l
}

/**
Double_variable_info {
    u1 tag = ITEM_Double; // 3
}
*/
type DoubleVariableInfo struct {
	Tag uint8
}

func (d DoubleVariableInfo) ReadVerficationTypeInfo(reader class_file_commons.Reader) VerificationTypeInfo {
	return d
}
