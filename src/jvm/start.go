package jvm

import (
	"errors"
	"fmt"
	"jvm/class/attribute"
	"jvm/class/class_file_commons"
	"jvm/class/constant_pool"
	fields "jvm/class/fields"
	"jvm/class/interfaces"
	methods "jvm/class/methods"
	"jvm/classloader"
	"jvm/commons"
)

/**
开始类加载，它有3个主要阶段：
	1. 加载
	2. 解析
	3. 初始化
检查参数、构建类加载器都是加载阶段、读取类的二进制码都属于加载阶段，目前正在实现
*/
func StartJVM() {
	options := new(commons.Cmd).ParseCmd()
	fmt.Printf("bootclasspath: %s\nextclasspath: %s\nclasspath: %s \nclass: %s\nargs:%v\n",
		options.BootClassPath, options.ExtClassPath, options.ClassPath, options.ClassName, options.Args)

	// TODO 检查启动参数[classloader, class]是否合法
	// checkOptionPoint(options);

	// TODO 1. 加载阶段
	var (
		stream     []byte
		err        error
		baseloader classloader.ClassLoader
		apploader  classloader.BaseLoader
		classPath  classloader.ClassPath = classloader.ClassPath{}
	)

	err = classPath.InitClassPath(*options)
	if err != nil {
		panic("classpath initializer error.")
	}

	stream, baseloader, err, apploader = classPath.UserClassLoader.ParentLoader(options.ClassName)
	if err != nil {
		panic(errors.New(fmt.Sprintf("ClassName=[%s] loader error, ClassLoader=[%s]\n", options.ClassName, apploader.GetName())))
	}

	// TODO 2. 进入解析阶段
	fmt.Println("data: ", stream)
	fmt.Printf("baseloader: %s\n", baseloader.ToString())
	fmt.Printf("apploader: %s\n", apploader.GetName())

	reader := class_file_commons.ClassFileReader{ByteStream: &stream}
	fmt.Printf("magic: %X\n", reader.ReadUint32())
	fmt.Printf("minor: %d\n", reader.ReadUint16())
	fmt.Printf("major: %d\n", reader.ReadUint16())

	// 解析常量池
	cr := constant_pool.ConstantPoolStructReader{&reader}
	cp_count, cp_infos := cr.ReadConstantPoolInfos()
	fmt.Println("constant_pool_count: ", cp_count)
	fmt.Println("constant_pool: ", cp_infos)

	fmt.Printf("access_flags: %X\n", reader.ReadUint16())
	fmt.Printf("this_class: %X\n", reader.ReadUint16())
	fmt.Printf("super_class: %X\n", reader.ReadUint16())

	// 解析接口表
	ir := interfaces.InterfacesStructReader{&cr}
	i_count, i_infos := ir.ReadInterfacesInfos()
	fmt.Println("interfaces_count: ", i_count)
	fmt.Println("interfaces: ", i_infos)

	// 解析字段表
	fr := fields.FieldsStructReader{&ir}
	f_count, f_infos := fr.ReadFieldsInfos()
	fmt.Println("fields_count: ", f_count)
	fmt.Println("fields: ", f_infos)

	// 解析方法表
	mr := methods.MethodsStructReader{&fr}
	m_count, m_infos := mr.ReadMethodsInfos()
	fmt.Println("methods_count: ", m_count)
	fmt.Println("methods: ", m_infos)

	// 解析类参数表
	ar := attribute.AttributesStructReader{&mr}
	a_count, a_infos := ar.ReadAttributeInfos()
	fmt.Println("attributes_count: ", a_count)
	fmt.Println("attributes: ", a_infos)
}
