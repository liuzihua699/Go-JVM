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
	CONSTANT_VALUE                           = "ConstantValue"
	CODE                                     = "Code"
	STACK_MAP_TABLE                          = "StackMapTable"
	EXCEPTIONS                               = "Exceptions"
	INNER_CLASSES                            = "InnerClasses"
	ENCLOSING_METHOD                         = "EnclosingMethod"
	SYNTHETIC                                = "Synthetic"
	SIGNATURE                                = "Signature"
	SOURCE_FILE                              = "SourceFile"
	SOURCE_DEBUG_EXTENSION                   = "SourceDebugExtension"
	LINE_NUMBER_TABLE                        = "LineNumberTable"
	LOCAL_VARIABLE_TABLE                     = "LocalVariableTable"
	LOCAL_VARIABLE_TYPE_TABLE                = "LocalVariableTypeTable"
	DEPRECATED                               = "Deprecated"
	RUNTIME_VISIBLE_ANNOATIONS               = "RuntimeVisibleAnnotations"
	RUNTIME_INVISIBLE_ANNOTATIONS            = "RuntimeInvisibleAnnotations"
	RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS    = "RuntimeVisibleParameterAnnotations"
	RUNTIME_INVISIBLE_PARAMETER_ANNOTAATIONS = "RuntimeInvisibleParameterAnnotations"
	RUNTIME_VISIBLE_TYPE_ANNOTATIONS         = "RuntimeVisibleTypeAnnotations"
	RUNTIME_INVISIBLE_TYPE_ANNOTATIONS       = "RuntimeInvisibleTypeAnnotations"
	ANNOTATION_DEFAULT                       = "AnnotationDefault"
	BOOTSTRAP_METHODS                        = "BootstrapMethods"
	METHOD_PARAMETERS                        = "MethodParameters"
)

func newAttributeInfo(attrNameIndex uint16, attrLength uint32) AttrInfo {
	name := constant_pool.GetFileConstantPool().Utf8Map[attrNameIndex]
	switch name {
	case CONSTANT_VALUE:
		return &ConstantValue_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case SIGNATURE:
		return &Signature_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case CODE:
		return &Code_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case LINE_NUMBER_TABLE:
		return &LineNumberTable_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case STACK_MAP_TABLE:
		return &StackMapTable_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case EXCEPTIONS:
		return &Exceptions_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case DEPRECATED:
		return &Deprecated_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case RUNTIME_VISIBLE_ANNOATIONS:
		return &RuntimeVisibleAnnotations_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case SOURCE_FILE:
		return &SourceFile_attribute{
			AttributeNameIndex: attrNameIndex,
			AttributeLength:    attrLength,
		}
	case INNER_CLASSES:
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
