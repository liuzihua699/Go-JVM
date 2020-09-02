package main

import (
	"os"
)

type JVMOption struct {
	author  string // 作者
	version string // 版本号
	time    string // 发布日期
	mode    State  // 模式
}

type State string

const dev State = "dev"   // 开发状态
const prod State = "prod" // 发布状态
const test State = "mode" // 测试状态

var JRE = getDefaultClassPath()
var BOOTSTRAPE_CLASS_PATH = JRE + "\\jre\\lib"
var EXT_CLASS_PATH = BOOTSTRAPE_CLASS_PATH + "\\ext"
var USER_CLASS_PATH = "."

/**
global config
*/
func getJVMOptions() JVMOption {
	return JVMOption{
		author:  "zihua",
		version: "0.0.3",
		time:    "2020/9/2",
		mode:    dev,
	}
}

/**
get classloader for system env.
e.g JAVA_HOME/
*/
func getDefaultClassPath() string {
	return os.Getenv("JAVA_HOME")
}
