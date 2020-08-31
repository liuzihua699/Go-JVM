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
}

func (c *Cmd) parseCmd() *Cmd {
	var cmd = &Cmd{}

	flag.Usage = c.printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version")
	flag.StringVar(&cmd.classPath, "classpath", "", "classpath")
	flag.StringVar(&cmd.classPath, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()

	if cmd.versionFlag {
		fmt.Println("Version for 0.1.")
		os.Exit(0)
	} else if cmd.helpFlag {
		c.printUsage()
		os.Exit(0)
	}

	if len(args) <= 0 {
		c.printNoArgument()
		os.Exit(0)
	}

	cmd.class = args[0]
	cmd.args = args[1:]
	return cmd
}

func (c *Cmd) printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func (c *Cmd) printNoArgument() {
	fmt.Println("You has no argument, please inpu [-help] or [-?] watch command line.")
}

/**
start the jvm
*/
func (c *Cmd) startJVM() {
	var cmd = c.parseCmd()
	fmt.Printf("classpath: %s \nclass: %s\nargs:%v\n",
		cmd.classPath, cmd.class, cmd.args)
}
