# 1. 设计命令行启动选项
2020年8月31日17:32:35

支持命令行 | 描述 | 是否实现
---|---|---
-help 或 -? | 输出帮助信息| √
-version 或 -v | 输出版本信息| √
-mode 或 -m| 输出当前模式| √
-author | 输出作者信息| √
-global_config | 输出全局配置| √
-Xbootclasspath | 指定启动类路径 | √
-Xextclasspath | 指定扩展类路径 | √
-cp %classpath% %class% %args% 或 -classpath| 加载字节码文件| √
%class% | 以默认的类路径加载字节码文件 | √

设计完成

（对于类路径的配置，暂不支持环境变量，仅支持运行的时候指定配置项）


# 2. 类加载器设计
## 2.1 classpath

