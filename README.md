# 1. 设计命令行启动选项
2020年8月31日17:32:35

我希望这样使用：
```cmd
./jvm -version 输出版本信息
./jvm -help 或 ? 输出帮助信息
./jvm -author

# 从classpath中运行一个类
./jvm -cp %classpath% %class% %args%
```
设计完毕