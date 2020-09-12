package constant_pool

import (
	"errors"
	"jvm/class/class_file_commons"
)

/*
e.g constant_pool format:
	cp_info {
		u1 tag;
		u1 info[];
	}
*/

/**
see https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.1
about 4.4 The Constant Pool
*/

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

var CONSTANT_POOL_TAG_NAME_MAP = map[uint8]string{
	CONSTANT_Class:              "CONSTANT_Class",
	CONSTANT_Fieldref:           "CONSTANT_Fieldref",
	CONSTANT_Methodref:          "CONSTANT_Methodref",
	CONSTANT_InterfaceMethodref: "CONSTANT_InterfaceMethodref",
	CONSTANT_String:             "CONSTANT_String",
	CONSTANT_Integer:            "CONSTANT_Integer",
	CONSTANT_Float:              "CONSTANT_Float",
	CONSTANT_Long:               "CONSTANT_Long",
	CONSTANT_Double:             "CONSTANT_Double",
	CONSTANT_NameAndType:        "CONSTANT_NameAndType",
	CONSTANT_Utf8:               "CONSTANT_Utf8",
	CONSTANT_MethodHandle:       "CONSTANT_MethodHandle",
	CONSTANT_MethodType:         "CONSTANT_MethodType",
	CONSTANT_InvokeDynamic:      "CONSTANT_InvokeDynamic",
}

//var CONSTANT_POOL_MAP = map[uint8]ConstantPoolInfo{
//	CONSTANT_Class:              &ConstantClassInfo{TagInfo: TagInfo{Tag: CONSTANT_Class}},
//	CONSTANT_Fieldref:           &ConstantFieldrefInfo{TagInfo: TagInfo{Tag: CONSTANT_Fieldref}},
//	CONSTANT_Methodref:          &ConstantMethodrefInfo{TagInfo: TagInfo{Tag: CONSTANT_Methodref}},
//	CONSTANT_InterfaceMethodref: &ConstantInterfacerefInfo{TagInfo: TagInfo{Tag: CONSTANT_InterfaceMethodref}},
//	CONSTANT_String:             &ConstantStringInfo{TagInfo: TagInfo{Tag: CONSTANT_String}},
//	CONSTANT_Integer:            &ConstantIntegerInfo{TagInfo: TagInfo{Tag: CONSTANT_Integer}},
//	CONSTANT_Float:              &ConstantFloatInfo{TagInfo: TagInfo{Tag: CONSTANT_Float}},
//	CONSTANT_Long:               &ConstantLongInfo{TagInfo: TagInfo{Tag: CONSTANT_Long}},
//	CONSTANT_Double:             &ConstantDoubleInfo{TagInfo: TagInfo{Tag: CONSTANT_Double}},
//	CONSTANT_NameAndType:        &ConstantNameAndTypeInfo{TagInfo: TagInfo{Tag: CONSTANT_NameAndType}},
//	CONSTANT_Utf8:               &ConstantUtf8Info{TagInfo: TagInfo{Tag: CONSTANT_Utf8}},
//	CONSTANT_MethodHandle:       &ConstantMethodHandleInfo{TagInfo: TagInfo{Tag: CONSTANT_MethodHandle}},
//	CONSTANT_MethodType:         &ConstantMethodTypeInfo{TagInfo: TagInfo{Tag: CONSTANT_MethodType}},
//	CONSTANT_InvokeDynamic:      &ConstantInvokeDynamicInfo{TagInfo: TagInfo{Tag: CONSTANT_InvokeDynamic}},
//}

type ConstantPoolInfo interface {
	ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo
	GetTag() uint8
	GetTagName() string
}

type ConstantPool struct {
	ConstantItemInfos []ConstantPoolInfo
	Utf8Map           map[uint16]string
}

func newConstantInfo(tag uint8) ConstantPoolInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{TagInfo: TagInfo{Tag: CONSTANT_Class}}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{TagInfo: TagInfo{Tag: CONSTANT_Fieldref}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{TagInfo: TagInfo{Tag: CONSTANT_Methodref}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfacerefInfo{TagInfo: TagInfo{Tag: CONSTANT_InterfaceMethodref}}
	case CONSTANT_String:
		return &ConstantStringInfo{TagInfo: TagInfo{Tag: CONSTANT_String}}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{TagInfo: TagInfo{Tag: CONSTANT_Integer}}
	case CONSTANT_Float:
		return &ConstantFloatInfo{TagInfo: TagInfo{Tag: CONSTANT_Float}}
	case CONSTANT_Long:
		return &ConstantLongInfo{TagInfo: TagInfo{Tag: CONSTANT_Long}}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{TagInfo: TagInfo{Tag: CONSTANT_Double}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{TagInfo: TagInfo{Tag: CONSTANT_NameAndType}}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{TagInfo: TagInfo{Tag: CONSTANT_Utf8}}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{TagInfo: TagInfo{Tag: CONSTANT_MethodHandle}}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{TagInfo: TagInfo{Tag: CONSTANT_MethodType}}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{TagInfo: TagInfo{Tag: CONSTANT_InvokeDynamic}}
	default:
		panic(errors.New("not fount tag handle error: invalid type to async of " + string(tag)))
		return nil
	}
}

var fileConstantPool *ConstantPool

func setFileConstantPool(pool *ConstantPool) {
	fileConstantPool = pool
}

func GetFileConstantPool() *ConstantPool {
	if fileConstantPool == nil {
		panic(errors.New("runtime error: class file constant pool has not initializer."))
	}
	return fileConstantPool
}
