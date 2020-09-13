package attribute

import (
	"errors"
	"jvm/class/class_file_commons"
	"jvm/class/constant_pool"
)

type Attributes struct {
	Attributes []AttributeInfo
}

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

func newAttributeInfo(info AttributeInfo) AttrInfo {
	name := constant_pool.GetFileConstantPool().Utf8Map[info.AttributeNameIndex]
	switch name {
	case ConstantValue:
		return &ConstantValue_attribute{AttributeInfo: info}
	case Signature:
		return &Signature_attribute{AttributeInfo: info}
	case Code:
		return &Code_attribute{AttributeInfo: info}
	case LineNumberTable:
		return &LineNumberTable_attribute{AttributeInfo: info}
	case StackMapTable:
		return &StackMapTable_attribute{AttributeInfo: info}
	case Exceptions:
		return &Exceptions_attribute{AttributeInfo: info}
	case Deprecated:
		return &Deprecated_attribute{AttributeInfo: info}
	case RuntimeVisibleAnnotations:
		return &RuntimeVisibleAnnotations_attribute{AttributeInfo: info}
	case SourceFile:
		return &SourceFile_attribute{AttributeInfo: info}
	case InnerClasses:
		return &InnerClasses_attribute{AttributeInfo: info}
	default:
		panic(errors.New("runtime error: can't find implemented attribute " + name))
		return nil
	}
}
