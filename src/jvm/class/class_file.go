package class

import (
	"jvm/class/constant_pool"
	fields "jvm/class/fields"
	methods "jvm/class/methods"
)

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	magic             uint32
	minjorVersion     uint16
	majorVersion      uint16
	constantPoolCount uint16
	constantPool      *constant_pool.ConstantPool
	accessFlags       uint16
	thisClass         uint16
	superClass        uint16
	interfacesCount   uint16
	interfaces        []uint16
	fieldsCount       uint16
	fields            *fields.Fields
	methodsCount      uint16
	methods           *[]methods.Methods
	attributesCount   uint16
	//attributes 			[]*AttributeInfo
}

// 字节码校验
/*func (c ClassFile) check() bool {
	return true
}*/

/*func (c ClassFile) Parse(byteStream []byte) (*ClassFile, error) {

}*/
