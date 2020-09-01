package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag      bool
	versionFlag   bool
	classPath     string
	class         string
	args          []string
	authorFlag    bool
	modeFlag      bool
	globalFlag    bool
	bootClassPath string
	extClassPath  string
}

// inject global config
var GLOBAL_CONFIG JVMOption = getJVMOptions()

func (c *Cmd) parseCmd() {

	flag.Usage = c.printUsage
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")

	flag.BoolVar(&c.authorFlag, "author", false, "please author")
	flag.BoolVar(&c.versionFlag, "version", false, "print version")
	flag.BoolVar(&c.versionFlag, "v", false, "print version")
	flag.BoolVar(&c.modeFlag, "mode", false, "print current mode")
	flag.BoolVar(&c.modeFlag, "m", false, "print current mode")
	flag.BoolVar(&c.globalFlag, "global_config", false, " print global config")

	flag.StringVar(&c.classPath, "classpath", USER_CLASS_PATH, "classpath")
	flag.StringVar(&c.classPath, "cp", USER_CLASS_PATH, "classpath")
	flag.StringVar(&c.bootClassPath, "Xbootclasspath", BOOTSTRAPE_CLASS_PATH, "print bootstrape classpath")
	flag.StringVar(&c.extClassPath, "Xextclasspath", EXT_CLASS_PATH, "print extension classpath")
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
}

func (c *Cmd) printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func (c *Cmd) printNoArgument() {
	fmt.Println("You has no argument, please input [-help] or [-?] watch help.")
}

/**
start the jvm
*/
func (c *Cmd) startJVM() {

	c.parseCmd()

	fmt.Printf("bootclasspath: %s\nextclasspath: %s\nclasspath: %s \nclass: %s\nargs:%v\n",
		c.bootClassPath, c.extClassPath, c.classPath, c.class, c.args)

	// TODO. 检查启动参数[classpath, class]是否合法
	// checkOptionPoint(c);

}
