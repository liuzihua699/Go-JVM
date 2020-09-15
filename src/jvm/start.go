package jvm

import (
	"errors"
	"fmt"
	"jvm/class"
	"jvm/classloader"
	"jvm/commons"
)

/**
开始类加载，它有3个主要阶段：
	1. 加载
	2. 解析
	3. 初始化
检查参数、构建类加载器都是加载阶段、读取类的二进制码都属于加载阶段，目前正在实现

检查参数，根据参数类型执行不同的行为：
	1. 如果指定了-p/-print，那么只打印类信息
	2. 否则会运行指定类的main方法
*/
func StartJVM() {
	options := new(commons.Cmd).ParseCmd()
	fmt.Printf("bootclasspath: %s\nextclasspath: %s\nclasspath: %s \nclass: %s\nargs:%v\n",
		options.BootClassPath, options.ExtClassPath, options.ClassPath, options.ClassName, options.Args)

	// TODO 检查启动参数[classloader, class]是否合法
	// checkOptionPoint(options);

	// 1. 加载阶段
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

	// 2. 解析字节码
	class, errs := new(class.ClassFile).Parse(stream)
	if len(errs) != 0 {
		panic(errs)
	}

	// 使用策略模式
	if options.PrintFlag {
		fmt.Println("BaseClassLoader: " + baseloader.ToString())
		fmt.Println("AppClassLoader: " + apploader.GetName())
		class.PrintClassInfo(!options.PrintModeAll)
	} else {
		// 运行main函数
	}
}
