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
DirClassLoader功能就是加载一个路径下的类

使用CreateDirLoader会定位到一个路径，然后传递的className参数有两种情况：
1. 全限定名(路径+类名)
2. 单类名

比如我在以`T:\\jvm-test`作为路径，以`HelloWorld`作为`className`，那么就会加载`T:\\jvm-test\\HelloWorld.class`；以`com.zihua.HelloWorld`作为`className`则会加载`T:\\jvm-test\\com\\zihua\\HelloWorld.class`。

具体实现见代码：Go-JVM/src/jvm/classloader/loader_dir.go

测试代码：Go-JVM/src/jvm/classloader/loader_dir_test.go

我的程序通过了以下单元测试：

```go
{
	name:    "1",
	fields:  fields{absDir: "T:\\jvm-test"},
	args:    args{className: "HelloWorld.class"},
},
{
	name:    "2",
	fields:  fields{absDir: "T:\\jvm-test"},
	args:    args{className: "HelloWorld"},
},
{
	name:    "3",
	fields:  fields{absDir: "T:\\jvm-test"},
	args:    args{className: "com.zihua.HelloWorld"},
},
{
	name:    "4",
	fields:  fields{absDir: "T:\\jvm-test"},
	args:    args{className: "HelloWorld.txt"},
},
```
### ZipClassLoader的设计
ZipClassLoader的作用是加载压缩文件，后缀为.zip或.jar，不区分大小写

创建该类加载器必须传递一个压缩文件的路径(不区分相对或绝对路径)，如果文件不存在则抛出异常。该类加载器和DirClassLoader的不同在于，DirClassLoader传递目录即可加载，而ZipClassLoader需要传递压缩文件。

具体实现见代码：Go-JVM/src/jvm/classloader/loader_zip.go

测试代码：Go-JVM/src/jvm/classloader/loader_zip_test.go


我的程序通过了如下单元测试：

```go
{
	name:    "1",
	path:    "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
	args: args{className: "java.lang.Class"},
},
{
	name:    "2",
	path:    "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
	args: args{className: "java.lang.String"},
},
{
	name:    "3",
	path:    "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\aaa.jar",
	args: args{className: "java.lang.String"},
},
{
	name:    "4",
	path:    "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
	args: args{className: "java.lang.zihua"},
},
```

### ComClassLoader的设计
ComClassLoader的功能就是在多个路径中加载一个类并返回，这里我不考虑用户故意构造不合法或不合逻辑的路径。

该类加载器逻辑上是DirClassLoader和ZipClassLoader的组合，允许用户传递`dir1;zip1;dir2;zip2`多个路径，用户需以系统分隔符分割。

我实现的思路就是分割路径，删除掉非法路径，然后对剩余路径创建ClassLoader，最后顺序调用LoadClass加载类即可。

具体实现见代码：Go-JVM/src/jvm/classloader/loader_composite.go

测试代码：Go-JVM/src/jvm/classloader/loader_composite_test.go

我的程序通过了如下单元测试：

```go
{
	name:    "1",
	args:    args{className: "com.zihua.HelloWorld"},
	pathList: "a;b;c;T:\\\\jvm-test;d",
},
{
	name:    "2",
	args:    args{className: "HelloWorld"},
	pathList: "a;b;c;T:\\\\jvm-test;d",
},
{
	name:    "3",
	args:    args{className: "java.lang.String"},
	pathList: "a;b;c;T:\\\\jvm-test;d",
},
{
	name:    "4",
	args:    args{className: "java.lang.String"},
	pathList: "a;b;c;T:\\\\jvm-test;d;C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
},
```

### WildcardClassLoader的设计
WildcardClassLoader会扫描目录并加载类。

目录下会有：
1. 子目录
2. .class文件
3. .jar或.zip文件


首先我会扫描这个目录下所有的目录和压缩文件并归档，先判断目录是否存在这个类，如果存在就先加载；然后按照顺序遍历jar包，如果加载到了类就返回；最后加载zip文件。都找不到就返回错误。

具体实现见代码：Go-JVM/src/jvm/classloader/loader_wildcard.go

测试代码：Go-JVM/src/jvm/classloader/loader_wildcard_test.go

我的程序通过了如下单元测试：

```go
{
	name: "1",
	args: args{className: "java.lang.Class"},
	path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\*",
},
{
	name: "2",
	args: args{className: "java.lang.Class"},
	path: "T:\\jvm-test\\*",
},
{
	name: "3",
	args: args{className: "com.zihua.HelloWorld"},
	path: "T:\\jvm-test\\*",
},
{
	name: "4",
	args: args{className: "java.lang.Class"},
	path: "T:\\jvm-test\\*",
},

```

## 2.2 加载基础类(classpath)
之前提到过，遍历路径由启动类路径、扩展类路径和用户类路径组成，如果我想根据某个类的全限定名加载这个类，那么需要通配符加载这些路径下的所有的路径，并判断是否存在基础类，比如java/lang/String。

## 2.3 双亲委派机制的实现？
BootClassLoader->ExtClassLoader->UserClassLoader

优先加载Boot，其次Ext，最后User：

```go
stream, baseloader, err = loader.BootClassLoader.LoadClass(options.class)
if err != nil {
	stream, baseloader, err = loader.ExtClassLoader.LoadClass(options.class)
	if err != nil {
		stream, baseloader, err = loader.UserClassLoader.LoadClass(options.class)
		if err != nil {
			panic(err)
		}
	}
}
```

这种方式的缺陷很明显：不支持自定义类加载器器

JVM官方的实现是用递归的方式，但我这个是属于1.0版本，所以先完成再优化。