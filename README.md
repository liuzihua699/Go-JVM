# 1. 启动选项
（对于类路径的配置，暂不支持环境变量，仅支持运行的时候指定配置项）

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

# 2. 类加载功能的设计
## 2.1 类加载器的设计
### 参数规范
java定义的类路径如下：
- 启动类路径：JAVA_HOME/jre/lib
- 扩展类路径：JAVA_HOME/jre/lib/ext
- 用户类路径：-cp指定的路径，不指定则为"."，代表当前目录

一般情况下-cp指定后会从指定的classpath加载class，如果不指定-cp那么就以JAVA_HOME/lib作为类路径加载class，如果没有配置JAVA_HOME则必须指定一个启动类路径(-Xbootclasspath)，所以命令行的启动需要把这个部分写清晰。

我需要做的也是按照这个类路径的标准去加载类环境，然后再加载类。流程为：
1. 加载类环境
2. 加载类
3. 从入口函数开始顺序执行

启动类和扩展类路径为class执行的基础环境，比如Object、String等类都是在这些基础类路径下的，用户类路径可选，不指定则会认为锁用到所有的类都处于基础环境下，如果不是则会报错。

### 类加载器接口定义
code in /src/classloader/*

目前接口暂时如下简单的定义：

```go
/**
  From 《Java Virtual Machine Specification》, I implement 4 type class-loader,
  They generate the corresponding class loader based on the path:
	1. DirClassLoader		e.g com/zihua/user/
	2. ZipClassLoader		e.g %JAVA_HOME%/jre/lib/rt.jar
	3. WildcardClassLoader	e.g com/zihua/*
	4. ComClassLoader		e.g com/zihua/system/;com/zihua/user/
*/
type ClassLoader interface {
	LoadClass(className string) ([]byte, ClassLoader, error)
	ToString() string
}
```

ClassLoader接受一个类的全限定名，然后加载返回这个类的字节码流。

对ClassLoader我定义了4种实现：
1. DirClassLoader：目录类加载器
2. ZipClassLoader：压缩文件类加载器
3. WildcardClassLoader：通配符路径类加载器
4. ComClassLoader：组合路径类加载器

比如我想加载 com/zihua/目录下的User.class类，那么会匹配到DirClassLoader然后加载。

对于一个类的初始化，我们首先得加载这个类的基类(java.lanng.Object)，还有很多常用的比如String也是必须优先加载，他们存在于JAVA_HOME/jre/lib/rt.jar文件中，这里就会匹配到压缩文件类加载器。

他们的拓扑关系如下：
1. dir_loader
2. com_loader->dir_loader
3. zip_loader
4. wild_loader->(zip_loader | dir_loader)

先设计dir_loader，然后再是com_loader和zip_loader，最后是wild_loader

### DirClassLoader的设计
### ZipClassLoader的设计
### ComClassLoader的设计
### WildcardClassLoader的设计

## 2.2 加载基础类(classpath)
之前提到过，遍历路径由启动类路径、扩展类路径和用户类路径组成，如果我想根据某个类的全限定名加载这个类，那么需要通配符加载这些路径下的所有的路径，并判断是否存在基础类，比如java/lang/String。

## 2.3 双亲委派机制的实现？