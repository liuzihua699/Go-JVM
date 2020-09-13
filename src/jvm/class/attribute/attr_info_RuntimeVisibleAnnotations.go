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
	AttributeInfo
	NumAnnotations uint16
	Annoations     []Annotation
}

func (r *RuntimeVisibleAnnotations_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	r.NumAnnotations = reader.ReadUint16()
	size := r.NumAnnotations
	for i := 0; i < int(size); i++ {
		r.Annoations = append(r.Annoations, new(Annotation).readAnnoation(reader))
	}
	return r
}

type Annotation struct {
	TypeIndex            uint16
	NumElementValuePairs uint16
	ElementValuePairs    []ElementValuePair
}

type ElementValuePair struct {
	ElementNameIndex uint16
	Value            ElementValue
}

type ElementValue struct {
	Tag   uint8
	Value UnionElementValue
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

func (e *ElementValuePair) readElementPair(reader class_file_commons.Reader) ElementValuePair {
	e.ElementNameIndex = reader.ReadUint16()
	e.Value = new(ElementValue).readElementValue(reader)
	return *e
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

// TODO too ugly code, if codereview that neet update ASCII to typename constant, e.g: 'B'=>BYTE
func newUnionElementValue(tag uint8) UnionElementValue {
	switch tag {
	case 'B':
		return &constValueIndex{Tag: tag}
	case 'C':
		return &constValueIndex{Tag: tag}
	case 'D':
		return &constValueIndex{Tag: tag}
	case 'F':
		return &constValueIndex{Tag: tag}
	case 'I':
		return &constValueIndex{Tag: tag}
	case 'J':
		return &constValueIndex{Tag: tag}
	case 'S':
		return &constValueIndex{Tag: tag}
	case 'Z':
		return &constValueIndex{Tag: tag}
	case 's':
		return &constValueIndex{Tag: tag}
	case 'e':
		return &enumValueIndex{Tag: tag}
	case 'c':
		return &classInfoIndex{Tag: tag}
	case '@':
		return &annotationValue{Tag: tag}
	case '[':
		return &arrayValue{Tag: tag}
	}
	panic(errors.New("runtime error: not found this tag, please watch document" + string(tag)))
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
