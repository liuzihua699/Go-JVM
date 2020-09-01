package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	classPath   string
	class       string
	args        []string
	authorFlag  bool
	modeFlag    bool
	globalFlag  bool
}

// inject global config
var global_config JVMOption = getJVMOptions()

func (c *Cmd) parseCmd() {

	flag.Usage = c.printUsage
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")
	flag.BoolVar(&c.authorFlag, "author", false, "please author")
	flag.BoolVar(&c.versionFlag, "version", false, "print version")
	flag.BoolVar(&c.versionFlag, "v", false, "print version")
	flag.StringVar(&c.classPath, "classpath", "", "classpath")
	flag.StringVar(&c.classPath, "cp", "", "classpath")
	flag.BoolVar(&c.modeFlag, "mode", false, "print current mode")
	flag.BoolVar(&c.modeFlag, "m", false, "print current mode")
	flag.BoolVar(&c.globalFlag, "global_config", false, " print global config")
	flag.Parse()
	args := flag.Args()

	if c.versionFlag {
		fmt.Printf("Version for %s\n", global_config.version)
		os.Exit(0)
	} else if c.helpFlag {
		c.printUsage()
		os.Exit(0)
	} else if c.authorFlag {
		fmt.Printf("author: %s\n", global_config.author)
		os.Exit(0)
	} else if c.modeFlag {
		fmt.Printf("mode: %s\n", global_config.mode)
		os.Exit(0)
	} else if c.globalFlag {
		fmt.Printf("author: %s\n", global_config.author)
		fmt.Printf("Version for %s\n", global_config.version)
		fmt.Printf("mode: %s\n", global_config.mode)
		fmt.Printf("time: %s\n", global_config.time)
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

	fmt.Printf("classpath: %s \nclass: %s\nargs:%v\n",
		c.classPath, c.class, c.args)

	// 检查JVM启动参数
	// checkOption();
}
