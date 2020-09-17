package attribute

import (
	"errors"
	"jvm/class/class_file_commons"
)

/**
RuntimeVisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}
*/
type RuntimeVisibleAnnotations_attribute struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumAnnotations     uint16
	Annoations         []Annotation
}

func (r *RuntimeVisibleAnnotations_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	r.NumAnnotations = reader.ReadUint16()
	size := r.NumAnnotations
	for i := 0; i < int(size); i++ {
		r.Annoations = append(r.Annoations, new(Annotation).readAnnoation(reader))
	}
	return r
}

/**
annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}
*/
type Annotation struct {
	TypeIndex            uint16
	NumElementValuePairs uint16
	ElementValuePairs    []ElementValuePair
}

func (a *Annotation) readAnnoation(reader class_file_commons.Reader) Annotation {
	a.TypeIndex = reader.ReadUint16()
	a.NumElementValuePairs = reader.ReadUint16()
	size := a.NumElementValuePairs
	for i := 0; i < int(size); i++ {
		a.ElementValuePairs = append(a.ElementValuePairs, new(ElementValuePair).readElementPair(reader))
	}
	return *a
}

/**
element_value_pair {
	u2 element_name_index
	element_value value
}
*/
type ElementValuePair struct {
	ElementNameIndex uint16
	Value            ElementValue
}

func (e *ElementValuePair) readElementPair(reader class_file_commons.Reader) ElementValuePair {
	e.ElementNameIndex = reader.ReadUint16()
	e.Value = new(ElementValue).readElementValue(reader)
	return *e
}

/**
element_value {
    u1 tag;
    union {
        u2 const_value_index;

        {   u2 type_name_index;
            u2 const_name_index;
        } enum_const_value;

        u2 class_info_index;

        annotation annotation_value;

        {   u2            num_values;
            element_value values[num_values];
        } array_value;
    } value;
}
*/
type ElementValue struct {
	Tag   uint8
	Value UnionElementValue
}

func (e *ElementValue) readElementValue(reader class_file_commons.Reader) ElementValue {
	e.Tag = reader.ReadUint8()
	tag := e.Tag
	e.Value = newUnionElementValue(tag).ReadUnionElementValue(reader)
	return *e
}

type UnionElementValue interface {
	ReadUnionElementValue(reader class_file_commons.Reader) UnionElementValue
}

const (
	TAG_BYTE            = 'B'
	TAG_CHAR            = 'C'
	TAG_DOUBLE          = 'D'
	TAG_FLOAG           = 'F'
	TAG_INT             = 'I'
	TAG_LONG            = 'J'
	TAG_SHORT           = 'S'
	TAG_BOOLEAN         = 'Z'
	TAG_STRING          = 's'
	TAG_ENUM_TYPE       = 'e'
	TAG_CLASS           = 'c'
	TAG_ANNOTATION_TYPE = '@'
	TAG_ARRAY_TYPE      = '['
)

// TODO too ugly code, if codereview that neet update ASCII to typename constant, e.g: 'B'=>BYTE
func newUnionElementValue(tag uint8) UnionElementValue {

	if tag == TAG_BYTE || tag == TAG_CHAR || tag == TAG_DOUBLE || tag == TAG_FLOAG ||
		tag == TAG_INT || tag == TAG_LONG || tag == TAG_SHORT || tag == TAG_BOOLEAN || tag == TAG_STRING {
		return &constValueIndex{Tag: tag}
	}

	switch tag {
	case TAG_ENUM_TYPE:
		return &enumValueIndex{Tag: tag}
	case TAG_CLASS:
		return &classInfoIndex{Tag: tag}
	case TAG_ANNOTATION_TYPE:
		return &annotationValue{Tag: tag}
	case TAG_ARRAY_TYPE:
		return &arrayValue{Tag: tag}
	}

	panic(errors.New("runtime error: not found this tag, please watch document , tag = " + string(tag)))
	return nil
}

type constValueIndex struct {
	Tag             uint8
	ConstValueIndex uint16
}

type enumValueIndex struct {
	Tag            uint8
	TypeNameIndex  uint16
	ConstNameIndex uint16
}

type classInfoIndex struct {
	Tag            uint8
	ClassInfoIndex uint16
}

type annotationValue struct {
	Tag            uint8
	AnnoationValue Annotation
}

type arrayValue struct {
	Tag      uint8
	NumValue uint16
	Values   []ElementValue
}

func (c *constValueIndex) ReadUnionElementValue(reader class_file_commons.Reader) UnionElementValue {
	c.ConstValueIndex = reader.ReadUint16()
	return c
}

func (e *enumValueIndex) ReadUnionElementValue(reader class_file_commons.Reader) UnionElementValue {
	e.TypeNameIndex = reader.ReadUint16()
	e.ConstNameIndex = reader.ReadUint16()
	return e
}

func (c *classInfoIndex) ReadUnionElementValue(reader class_file_commons.Reader) UnionElementValue {
	c.ClassInfoIndex = reader.ReadUint16()
	return c
}

func (a *annotationValue) ReadUnionElementValue(reader class_file_commons.Reader) UnionElementValue {
	a.AnnoationValue = new(Annotation).readAnnoation(reader)
	return a
}

func (a *arrayValue) ReadUnionElementValue(reader class_file_commons.Reader) UnionElementValue {
	a.NumValue = reader.ReadUint16()
	size := a.NumValue
	for i := 0; i < int(size); i++ {
		a.Values = append(a.Values, new(ElementValue).readElementValue(reader))
	}
	return a
}
