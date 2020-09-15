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

	// TODO 检查启动参数[classloader, class]是否合法
	// checkOptionPoint(options);

	var (
		stream     []byte
		err        error
		baseloader classloader.ClassLoader
		apploader  classloader.BaseLoader
		classPath  classloader.ClassPath = classloader.ClassPath{}
	)

	// initializer the classpath from user options.
	err = classPath.InitClassPath(*options)
	if err != nil {
		panic("classpath initializer error.")
	}

	// parent loader to load the class.
	stream, baseloader, err, apploader = classPath.UserClassLoader.ParentLoader(options.ClassName)
	if err != nil {
		panic(errors.New(fmt.Sprintf("ClassName=[%s] loader error, ClassLoader=[%s]\n", options.ClassName, apploader.GetName())))
	}

	// parse bytecode to class file format.
	class, errs := new(class.ClassFile).Parse(stream)
	if len(errs) != 0 {
		panic(errs)
	}

	// if print-only mode, that print bytecode class information.
	if options.PrintFlag {
		fmt.Printf("bootclasspath: %s\nextclasspath: %s\nclasspath: %s \nclass: %s\nargs:%v\n",
			options.BootClassPath, options.ExtClassPath, options.ClassPath, options.ClassName, options.Args)
		fmt.Println("BaseClassLoader: " + baseloader.ToString())
		fmt.Println("AppClassLoader: " + apploader.GetName())
		class.PrintClassInfo(!options.PrintModeAll)
	} else {
		// TODO 运行main函数
	}
}
