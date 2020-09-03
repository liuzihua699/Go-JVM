package commons

import (
	"os"
)

type JVMOption struct {
	Author     string // 作者
	Version    string // 版本号
	Start_time string // 开始日期
	Last_time  string // 发布日期
	Mode       State  // 模式
}

type State string

const dev State = "dev"   // 开发状态
const prod State = "prod" // 发布状态
const test State = "mode" // 测试状态

var JRE = getDefaultClassPath()
var BOOTSTRAPE_CLASS_PATH = JRE + "\\jre\\lib"
var EXT_CLASS_PATH = BOOTSTRAPE_CLASS_PATH + "\\ext"
var USER_CLASS_PATH = "."

// inject global config
var GLOBAL_CONFIG = GetJVMOptions()

/**
global config
*/
func GetJVMOptions() JVMOption {
	return JVMOption{
		Author:     "zihua",
		Version:    "0.0.4",
		Start_time: "2020/8/31",
		Last_time:  "2020/9/3",
		Mode:       dev,
	}
}

/**
get classloader for system env.
e.g JAVA_HOME/
*/
func getDefaultClassPath() string {
	return os.Getenv("JAVA_HOME")
}
