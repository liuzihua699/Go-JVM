package class

import (
	"errors"
	"fmt"
	"jvm/class/attribute"
	"jvm/class/class_file_commons"
	CP "jvm/class/constant_pool"
	fields "jvm/class/fields"
	"jvm/class/interfaces"
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
	constantPool      *CP.ConstantPool
	accessFlags       uint16
	thisClass         uint16
	superClass        uint16
	interfacesCount   uint16
	interfaces        []uint16
	fieldsCount       uint16
	fields            *fields.Fields
	methodsCount      uint16
	methods           *methods.Methods
	attributesCount   uint16
	attributes        *attribute.Attributes
}

// 字节码校验
/*func (c ClassFile) check() bool {
	return true
}*/

const (
	MAGIC      = 3405691582
	BASE_CLASS = 0
)

// TODO the class file to parse...
func (c *ClassFile) Parse(byteStream []byte) (ClassFile, []error) {
	errs := *new([]error)

	// 解析基本属性
	r := class_file_commons.ClassFileReader{ByteStream: &byteStream}
	c.magic = r.ReadUint32()
	c.minjorVersion = r.ReadUint16()
	c.majorVersion = r.ReadUint16()
	if c.magic^MAGIC != 0 {
		errs = append(errs, errors.New("error: reject parse the class type."))
	}

	// 解析文件常量池
	cr := CP.ConstantPoolStructReader{Reader: &r}
	cp_count, cp_infos := cr.ReadConstantPoolInfos()
	c.constantPoolCount = cp_count
	c.constantPool = cp_infos
	if cp_count == 0 {
		errs = append(errs, errors.New("error: constant pool parse error."))
	}

	// 解析基本属性
	c.accessFlags = r.ReadUint16()
	c.thisClass = r.ReadUint16()
	c.superClass = r.ReadUint16()

	// 解析接口表
	ir := interfaces.InterfacesStructReader{&r}
	i_count, i_infos := ir.ReadInterfacesInfos()
	c.interfacesCount = i_count
	c.interfaces = i_infos

	// 解析字段表
	fr := fields.FieldsStructReader{&r}
	f_count, f_infos := fr.ReadFieldsInfos()
	c.fieldsCount = f_count
	c.fields = f_infos

	// 解析方法表
	mr := methods.MethodsStructReader{&r}
	m_count, m_infos := mr.ReadMethodsInfos()
	c.methodsCount = m_count
	c.methods = m_infos

	// 解析参数表
	ar := attribute.AttributesStructReader{&r}
	a_count, a_infos := ar.ReadAttributeInfos()
	c.attributesCount = a_count
	c.attributes = a_infos

	return *c, errs
}

func (c ClassFile) PrintClassInfo(trunc bool) {
	fmt.Printf("magic: 0x%X\n", c.magic)
	fmt.Printf("minor version: %d\n", c.minjorVersion)
	fmt.Printf("major version: %d\n", c.majorVersion)
	fmt.Printf("constantpool count: %d\n", c.constantPoolCount)
	fmt.Printf("access flags: %b\n", c.accessFlags)

	Utf8Pool := c.constantPool.Utf8Map
	CpPool := c.constantPool.CpMap

	// 类信息
	nameIndex := CpPool[c.thisClass].(*CP.ConstantClassInfo).NameIndex
	fmt.Printf("this class: %s\n", Utf8Pool[nameIndex])
	if c.superClass != BASE_CLASS {
		nameIndex = CpPool[c.superClass].(*CP.ConstantClassInfo).NameIndex
		fmt.Printf("super class: %s\n", Utf8Pool[nameIndex])
	}

	// 接口表信息
	fmt.Println()
	fmt.Printf("interface count: %d\n", c.interfacesCount)
	if c.interfacesCount != 0 {
		fmt.Printf("interfaces: \n")
		ii := 0
		for i := range c.interfaces {
			index := CpPool[c.interfaces[i]].(*CP.ConstantClassInfo).NameIndex
			fmt.Printf("\t%s\n", Utf8Pool[index])
			ii++
			if trunc && ii >= 10 {
				fmt.Println("\t...")
				break
			}
		}
	}

	// 字段表信息
	fmt.Println()
	fmt.Printf("fields count: %d\n", c.fieldsCount)
	if c.fieldsCount != 0 {
		fmt.Printf("fields: \n")
		i := 0
		for _, field := range c.fields.MemberInfos {
			fmt.Printf("\t%s\n", field.GetName())
			i++
			if trunc && i >= 10 {
				fmt.Println("\t...")
				break
			}
		}
	}

	// 方法表信息
	fmt.Println()
	fmt.Printf("method count: %d\n", c.methodsCount)
	if c.methodsCount != 0 {
		fmt.Printf("methods: \n")
		i := 0
		for _, method := range c.methods.MemberInfos {
			fmt.Printf("\t%s%s\n", method.GetName(), method.GetDescripor())
			i++
			if trunc && i >= 10 {
				fmt.Println("\t...")
				break
			}
		}
	}

	// 不用打印属性表
	//fmt.Printf("attribute count: %d\n", c.attributesCount)
	//if c.attributesCount != 0 {
	//	fmt.Printf("attributes: ")
	//	i := 0
	//	for _, attr := range c.attributes.Attributes {
	//		fmt.Printf("\t%s\n", attr)
	//		if trunc && i >= 10 {
	//			fmt.Println("...")
	//			break
	//		}
	//	}
	//}
}
