package commons

import (
	"os"
)

type JVMOption struct {
	Author     string         // 作者
	Version    string         // 版本号
	Start_time string         // 开始日期
	Last_time  string         // 发布日期
	Mode       State          // 模式
	DevLog     map[string]Log // 版本日志
}

type Log struct {
	version  string
	describe string
	time     string
}

type State string

const dev State = "dev"   // 开发状态
const prod State = "prod" // 发布状态
const test State = "mode" // 测试状态

const SEPARATOR = string(os.PathSeparator)

var JDK = getDefaultClassPath()
var JRE = JDK + SEPARATOR + "jre"
var BOOTSTRAPE_CLASS_PATH = JRE + SEPARATOR + "lib"
var EXT_CLASS_PATH = BOOTSTRAPE_CLASS_PATH + SEPARATOR + "ext"
var USER_CLASS_PATH = "."

// inject global config
var GLOBAL_CONFIG = GetDevOptions()

/**
global config
*/
func GetDevOptions() JVMOption {
	return JVMOption{
		Author:     "zihua",
		Version:    "0.0.4",
		Start_time: "2020/8/31",
		Last_time:  "2020/9/10",
		Mode:       dev,
		DevLog: map[string]Log{
			"2020/8/31": {
				version:  "0.0.1",
				describe: "Cmd功能，支持部分参数",
			},
			"2020/9/1": {
				version:  "0.0.2",
				describe: "优化命令行选项",
			},
			"2020/9/2": {
				version:  "0.0.3",
				describe: "基础类加载器的完成",
			},
			"2020/9/3": {
				version:  "0.0.4",
				describe: "应用层类加载器的整合",
			},
			"2020/9/10": {
				version:  "0.0.5",
				describe: "9/9常量池解析完毕，9/10更正了一些代码逻辑使之正确解析常量池",
			},
			"2020/9/11": {
				version:  "0.0.6",
				describe: "完成字段表解析",
			},
			"2020/9/12": {
				version:  "0.0.7",
				describe: "完成属性表解析，支持大部分类型的属性解析(代码需要重构)",
			},
			"2020/9/13": {
				version:  "0.0.7",
				describe: "解析功能实现完毕",
			},
			"2020/9/14": {
				version:  "0.0.8",
				describe: "对attributes模块做code-review",
			},
			"2020/9/15": {
				version:  "0.1.0",
				describe: "添加-p参数输出类文件格式",
			},
		},
	}
}

/**
get classloader for system env.
e.g JAVA_HOME/
*/
func getDefaultClassPath() string {
	return os.Getenv("JAVA_HOME")
}
