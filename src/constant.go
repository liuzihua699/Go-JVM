package main

type JVMOption struct {
	author  string // 作者
	version string // 版本号
	time    string // 发布日期
	mode    State  // 模式
}

type State string

const (
	dev  State = "dev"  // 开发状态
	prod State = "prod" // 发布状态
	test State = "mode" // 测试状态
)

func getJVMOptions() JVMOption {
	return JVMOption{
		author:  "zihua",
		version: "0.0.2",
		time:    "2020/8/31",
		mode:    dev,
	}
}

/**
get classpath for system env.
*/
func getDefaultClassPath() {

}
