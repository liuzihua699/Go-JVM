package main

import (
	"flag"
	"fmt"
	"jvm/classloader"
	"os"
)

type Cmd struct {
	helpFlag      bool
	versionFlag   bool
	authorFlag    bool
	modeFlag      bool
	globalFlag    bool
	args          []string
	classPath     string
	class         string
	bootClassPath string
	extClassPath  string
}

// inject global config
var GLOBAL_CONFIG = getJVMOptions()

func (c *Cmd) parseCmd() *Cmd {

	flag.Usage = c.printUsage
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")

	flag.BoolVar(&c.authorFlag, "author", false, "please author")
	flag.BoolVar(&c.versionFlag, "version", false, "print version")
	flag.BoolVar(&c.versionFlag, "v", false, "print version")
	flag.BoolVar(&c.modeFlag, "mode", false, "print current mode")
	flag.BoolVar(&c.modeFlag, "m", false, "print current mode")
	flag.BoolVar(&c.globalFlag, "global_config", false, " print global config")

	flag.StringVar(&c.classPath, "classloader", USER_CLASS_PATH, "classloader")
	flag.StringVar(&c.classPath, "cp", USER_CLASS_PATH, "classloader")
	flag.StringVar(&c.bootClassPath, "Xbootclasspath", BOOTSTRAPE_CLASS_PATH, "print bootstrape classloader")
	flag.StringVar(&c.extClassPath, "Xextclasspath", EXT_CLASS_PATH, "print extension classloader")
	flag.Parse()
	args := flag.Args()

	if c.versionFlag {
		fmt.Printf("Version for %s\n", GLOBAL_CONFIG.version)
		os.Exit(0)
	} else if c.helpFlag {
		c.printUsage()
		os.Exit(0)
	} else if c.authorFlag {
		fmt.Printf("author: %s\n", GLOBAL_CONFIG.author)
		os.Exit(0)
	} else if c.modeFlag {
		fmt.Printf("mode: %s\n", GLOBAL_CONFIG.mode)
		os.Exit(0)
	} else if c.globalFlag {
		fmt.Printf("author: %s\n", GLOBAL_CONFIG.author)
		fmt.Printf("Version for %s\n", GLOBAL_CONFIG.version)
		fmt.Printf("mode: %s\n", GLOBAL_CONFIG.mode)
		fmt.Printf("time: %s\n", GLOBAL_CONFIG.time)
		os.Exit(0)
	}

	if len(args) <= 0 {
		c.printNoArgument()
		os.Exit(0)
	}

	c.class = args[0]
	c.args = args[1:]
	return c
}

func (c *Cmd) printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func (c *Cmd) printNoArgument() {
	fmt.Println("You has no argument, please input [-help] or [-?] watch help.")
}

/**
开始类加载，它有3个主要阶段：
	1. 加载
	2. 解析
	3. 初始化
检查参数、构建类加载器都是加载阶段、读取类的二进制码都属于加载阶段，目前正在实现
*/
func startJVM() {
	options := new(Cmd).parseCmd()
	fmt.Printf("bootclasspath: %s\nextclasspath: %s\nclasspath: %s \nclass: %s\nargs:%v\n",
		options.bootClassPath, options.extClassPath, options.classPath, options.class, options.args)

	// TODO 检查启动参数[classloader, class]是否合法
	// checkOptionPoint(options);

	// TODO 1.加载阶段
	loader := new(classloader.ClassPath)
	loader.BootClassLoader = classloader.CreateLoader(options.bootClassPath + "\\*")
	loader.ExtClassLoader = classloader.CreateLoader(options.extClassPath + "\\*")
	loader.UserClassLoader = classloader.CreateLoader(options.classPath)

	var (
		stream     []byte
		baseloader classloader.ClassLoader
		err        error
	)

	// 父委派机制的类加载实现
	stream, baseloader, err = loader.BootClassLoader.LoadClass(options.class)
	if err != nil {
		stream, baseloader, err = loader.ExtClassLoader.LoadClass(options.class)
		if err != nil {
			stream, baseloader, err = loader.UserClassLoader.LoadClass(options.class)
			if err != nil {
				panic(err)
			}
		}
	}

	// .TODO 2.进入解析阶段
	fmt.Println(stream)
	fmt.Println(baseloader.ToString())
}

func main() {
	startJVM()
}
