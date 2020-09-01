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
}

// inject global config
var global_config JVMOption = getJVMOptions()

func (c *Cmd) parseCmd() {

	flag.Usage = c.printUsage
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")
	flag.BoolVar(&c.authorFlag, "author", false, "please author")
	flag.BoolVar(&c.versionFlag, "version", false, "print version")
	flag.StringVar(&c.classPath, "classpath", "", "classpath")
	flag.StringVar(&c.classPath, "cp", "", "classpath")
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
