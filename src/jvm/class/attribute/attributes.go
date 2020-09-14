package attribute

import (
	"fmt"
	"jvm/class/class_file_commons"
	"jvm/class/constant_pool"
)

type Attributes struct {
	Attributes []AttrInfo
}

/**
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length]; // bytes array
}

see document https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.4
types number of 23.
*/
type AttrInfo interface {
	ReadAttrInfo(reader class_file_commons.Reader) AttrInfo
}

const (
	ConstantValue                        = "ConstantValue"
	Code                                 = "Code"
	StackMapTable                        = "StackMapTable"
	Exceptions                           = "Exceptions"
	InnerClasses                         = "InnerClasses"
	EnclosingMethod                      = "EnclosingMethod"
	Synthetic                            = "Synthetic"
	Signature                            = "Signature"
	SourceFile                           = "SourceFile"
	SourceDebugExtension                 = "SourceDebugExtension"
	LineNumberTable                      = "LineNumberTable"
	LocalVariableTable                   = "LocalVariableTable"
	LocalVariableTypeTable               = "LocalVariableTypeTable"
	Deprecated                           = "Deprecated"
	RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	RuntimeVisibleTypeAnnotations        = "RuntimeVisibleTypeAnnotations"
	RuntimeInvisibleTypeAnnotations      = "RuntimeInvisibleTypeAnnotations"
	AnnotationDefault                    = "AnnotationDefault"
	BootstrapMethods                     = "BootstrapMethods"
	MethodParameters                     = "MethodParameters"
)

func newAttributeInfo(attrNameIndex uint16, attrLength uint32) AttrInfo {
	name := constant_pool.GetFileConstantPool().Utf8Map[attrNameIndex]
	switch name {
	case ConstantValue:
		return &ConstantValue_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case Signature:
		return &Signature_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case Code:
		return &Code_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case LineNumberTable:
		return &LineNumberTable_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case StackMapTable:
		return &StackMapTable_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case Exceptions:
		return &Exceptions_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case Deprecated:
		return &Deprecated_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case RuntimeVisibleAnnotations:
		return &RuntimeVisibleAnnotations_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case SourceFile:
		return &SourceFile_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case InnerClasses:
		return &InnerClasses_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	default:
		//panic(errors.New("runtime error: can't find implemented attribute " + name))
		fmt.Errorf("runtime warning: can't find implemented attribute " + name)
		return &Default_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	}
}

func ReadAttributeInfo(reader class_file_commons.Reader) AttrInfo {
	attrNameIndex := reader.ReadUint16()
	attrLength := reader.ReadUint32()
	return newAttributeInfo(attrNameIndex, attrLength).ReadAttrInfo(reader)
}
