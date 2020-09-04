package jvm

import (
	"errors"
	"fmt"
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

	// TODO 1.加载阶段
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

	// .TODO 2.进入解析阶段
	fmt.Println("data: ", stream)
	fmt.Println("baseloader: ", baseloader.ToString())
	fmt.Println("apploader: ", apploader.GetName())

}
