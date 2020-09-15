package commons

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag       bool
	versionFlag    bool
	authorFlag     bool
	modeFlag       bool
	globalFlag     bool
	Args           []string
	ClassPath      string
	ClassName      string
	BootClassPath  string
	ExtClassPath   string
	JrePath        string
	PrintFlag      bool
	PrintModeAll   bool
	PrintModeTrunc bool
}

func (c *Cmd) ParseCmd() *Cmd {

	flag.Usage = c.printUsage
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")
	flag.BoolVar(&c.PrintFlag, "p", false, "")
	flag.BoolVar(&c.PrintFlag, "print", false, "")
	flag.BoolVar(&c.PrintModeAll, "no-trunc", false, "print all class information.")
	flag.BoolVar(&c.PrintModeTrunc, "trunc", true, "print trunc class information.")

	flag.BoolVar(&c.authorFlag, "author", false, "please author")
	flag.BoolVar(&c.versionFlag, "version", false, "print version")
	flag.BoolVar(&c.versionFlag, "v", false, "print version")
	flag.BoolVar(&c.modeFlag, "mode", false, "print current mode")
	flag.BoolVar(&c.modeFlag, "m", false, "print current mode")
	flag.BoolVar(&c.globalFlag, "global_config", false, " print global config")

	flag.StringVar(&c.ClassPath, "classpath", USER_CLASS_PATH, "classpath")
	flag.StringVar(&c.ClassPath, "cp", USER_CLASS_PATH, "classloader")
	flag.StringVar(&c.BootClassPath, "Xbootclasspath", BOOTSTRAPE_CLASS_PATH, "print bootstrape classloader")
	flag.StringVar(&c.ExtClassPath, "Xextclasspath", EXT_CLASS_PATH, "print extension classloader")
	flag.StringVar(&c.JrePath, "Xjre", JRE, "print jre path")
	flag.Parse()
	args := flag.Args()

	if c.versionFlag {
		fmt.Printf("Version for %s\n", GLOBAL_CONFIG.Version)
		os.Exit(0)
	} else if c.helpFlag {
		c.printUsage()
		os.Exit(0)
	} else if c.authorFlag {
		fmt.Printf("author: %s\n", GLOBAL_CONFIG.Author)
		os.Exit(0)
	} else if c.modeFlag {
		fmt.Printf("mode: %s\n", GLOBAL_CONFIG.Mode)
		os.Exit(0)
	} else if c.globalFlag {
		fmt.Printf("author: %s\n", GLOBAL_CONFIG.Author)
		fmt.Printf("Version for %s\n", GLOBAL_CONFIG.Version)
		fmt.Printf("mode: %s\n", GLOBAL_CONFIG.Mode)
		fmt.Printf("time: %s\n", GLOBAL_CONFIG.Start_time)
		os.Exit(0)
	}

	if len(args) <= 0 {
		c.printNoArgument()
		os.Exit(0)
	}

	c.ClassName = args[0]
	c.Args = args[1:]
	return c
}

func (c *Cmd) printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func (c *Cmd) printNoArgument() {
	fmt.Println("You has no argument, please input [-help] or [-?] watch help.")
}
