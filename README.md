# 项目地址
https://github.com/zihuaVeryGood/Go-JVM

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
-Xjre|指定jre目录|√
-cp %classpath% %class% %args% 或 -classpath| 加载字节码文件| √
%class% | 以默认的类路径加载字节码文件 | √

设计完成

(等到1.0版本开发完毕就会发布releases和tag)


新增功能
支持命令行 | 描述 | 是否实现
---|---|---
-p/-print [--all | --trunc] | 打印信息[详细打印 | 截取打印]| √


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
4. wild_loader->(zip_loader + dir_loader)

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

## 2.2 整体加载流程
遍历路径由启动类路径、扩展类路径和用户类路径组成，如果我想根据某个类的全限定名加载这个类，那么需要通配符加载这些路径下的所有的路径，并判断是否存在基础类，比如java/lang/String。

所以最终的加载公式为：

```
pathlist = composite(wildcard("%JAVA_HOME%/jre/lib"), wildcard("%JAVA_HOME%/jre/lib/ext"), dir("%classpath%"))
```

0.0.3版本的程序通过了如下测试(已在程序中打包)：


加载参数 | 测试描述 | 预期结果 | 是否通过
---|---|---|---
-cp T:\\jvm-test HelloWorld | 类放至boot、ext、user目录下 | 使用BootClassLoader加载该类 | √
-cp T:\\jvm-test HelloWorld | 类放至ext、user目录下 | 使用ExtClassLoader加载该类 | √
-cp T:\\jvm-test HelloWorld | 类放至user目录下 | 使用UserClassLoader加载该类 | √
-cp T:\\jvm-test java.lang.String | rt.jar放至boot、ext、user目录下 | 使用BootClassLoader加载该类 | √
-cp T:\\jvm-test java.lang.String | rt.jar放至ext、user目录下 | 使用ExtClassLoader加载该类 | √
-cp T:\\jvm-test java.lang.String | String.class放至user/java/lang目录下 | 使用UserClassLoader加载该类 | √


目前已知问题：
1. 加载java.lang.String时，由于ext在boot目录下，所以加载器还是boot加载器，但总体没有太大影响。解决方法是放弃使用通配符加载器而使用路径加载器
2. 如果不指定classpath不确定会不会有什么影响


## 2.3 双亲委派机制的实现？
启动顺序为：BootClassLoader->ExtClassLoader->UserClassLoader

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

JVM官方的实现是用递归的方式，所以我使用递归的写法：

```go
func (l BaseLoader) ParentLoader(className string) ([]byte, ClassLoader, error, BaseLoader) {
	if l.hasParent() {
		data, loader, err, ll := l.ParentClassLoader.ParentLoader(className)
		if err == nil {
			return data, loader, err, ll
		}
		a, b, c := l.loadClass(className)
		return a, b, c, l
	}
	a, b, c := l.loadClass(className)
	return a, b, c, l
}
```

只要存在父加载器就递归直至最顶层的加载器，然后一层一层的去loadClass，当顶层执行成功则会直接返回信息，执行失败则会回溯至下一层调用下一层的loadClass，如果失败则继续回溯直至成功。

# 3. 字节码的设计
## 3.1 字节码规范
在Oracle的官方《Java虚拟机规范》中描述了几点：

具体可以翻[原文档](https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html)Chapter4。

一个类文件由8位字节流组成，通过连续读取2、4和8个连续的8位字节来构造16位、32位和64位量，多字节数据存储以big-endian的方式存储。

官方定义ClassFile格式为：
```c
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
```

字段含义如下：

字段名 | 描述
---|---
magic | 为ClassFile提供标识，为固定值0xCAFEBABE
minor_version | ClassFile最小版本号，JVM版本低于minor则会抛异常
major_version | ClassFile最大版本号，JVM版本高于major则会抛异常
constant_pool_count | consstant_pool的数量
constant_pool | 用来表示各种字符串常量、类、接口名、字段名和其他常量。下标从0开始。
access_flags | 访问标志，标识类或接口的访问权限，它代表当前常量池中所有参数的flag(均为access_flags)，具体值请看Chapter4中的参数详解
this_class | 这个值必须在常量池constatn_pool中索引，而且必须为CONSTANT_Class_info结构体。这个字段代表当前类或接口的类文件。
super_class | 这个字段要么为0，要么为constatn_pool中的有效索引，不为0则必须为CONSTAN_Class_info结构体表示为当前类直接的父类。直接超类不允许被设置ACC_FINAL
interface_count | interface[]的数量
interface[] | 可以在contant_pool中找到索引


## 3.2 字节码定义与设计
构成class的单位是字节，可以把整个class文件当做字节流处理。

JVM规范中对字节码的定义是大端(big-endain)的方式，所以在go中定义如下接口读取字节流：

```go
type Reader interface {
	ReadUint8() 			 	uint8 		// 读取1字节，对应java中的bit
	ReadUint16() 			 	uint16 		// 读取2字节，对应java中的short
	ReadUint32() 			 	uint32 		// 读取4字节，对应java中的int
	ReadUint64() 			 	uint64 		// 读取8字节，对应java中的long
	ReadBytes(len uint32) 		[]byte 		// 读取定长字节数
	Clear()
}
```


## 3.3 常量池的定义与设计
`constant_pool`字段是由以下结构组成：
```
cp_info {
    u1 tag;
    u1 info[];
}
```
其中tag可以为以下的值，并表示其字面的含义：

Constant Type |	Value
---|---
CONSTANT_Class	|7
CONSTANT_Fieldref|	9
CONSTANT_Methodref	|10
CONSTANT_InterfaceMethodref	|11
CONSTANT_String	|8
CONSTANT_Integer|	3
CONSTANT_Float|	4
CONSTANT_Long|	5
CONSTANT_Double	|6
CONSTANT_NameAndType|	12
CONSTANT_Utf8|	1
CONSTANT_MethodHandle|	15
CONSTANT_MethodType	|16
CONSTANT_InvokeDynamic	|18


读取常量池时，会先读入1字节的`tag`，然后根据tag的值读取不同的结构体，最后组成的cp_info数组为常量池的组成。

由于类型太多(14个)，所以我不会全部放出来占位置，具体情况感兴趣的可以去[官网](https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.1)4.4开始查看常量池结构。

我的实现思路如下：
1. 定义tag结构体对应的handler，具体见(jvm/class/constant_pool/constant_pool.go)
2. 定义常量池reader，具体见(jvm/class/constant_pool/constant_pool_reader.go)
3. 定义tag处理器，具体见(jvm/class/constant_pool/constant_*_info.go)

在解析常量池的过程中遇到了问题，比如在解析java.lanng.String的时候，会提示异常，经过debug追踪，问题发生在无法解析tag，因为此时读取到的tag为0。所以找到的hadnle为nil，所以报错了。

经过一些人的提示，发现解析常量池需要注意2个点：
1. 常量池的实际大小比constant_pool_size要小1
2. Double和Long实际要占2个位置，也就是说如果遇到Double和Long的tag，它们的下一个位置应该为nil且解析的size总长也要-1


修改代码后我重新测试了如下类的解析，均成功：
1. java.lang.String
2. java.lang.Integer
3. java.lang.Class
4. java.lang.ClassLoader

## 3.4 接口表的定义与设计
接口表在官方定义中是一个u2数组，也就是无需什么特殊的数据结构来处理，使用uint16即可。

具体读取代码请参考(jvm/class/interfaces/interfaces_reader.go)

## 3.5 字段表的定义与设计
2字节表示字段表的数量，剩余为一个filed结构体

```c
field_info {
	u2             access_flags;
	u2             name_index;
	u2             descriptor_index;
	u2             attributes_count;
	attribute_info attributes[attributes_count];
}
```

实现代码请参考(jvm/class/fields/*.go)


其中attribute_info表示属性表，这部分请参考属性表的设计

## 3.6 方法表的定义与设计
代码参考(jvm/methods/*.go)

2字节表示方法表数量，剩余为一个method结构体：

```c
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
```

其中attribute_info表示属性表，这部分请参考属性表的设计




## 3.7 属性表的定义与设计

代码参考(jvm/attributes/*.go)


属性表结构体非常麻烦，官方对其大致划分如下：

```c
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
```

简单的说，我可以读取2字节的index，这个index在常量池中是一个utf8的索引，我们可以获取这个index对应的值，然后读取4字节长度，最后的info装的是length这么长的字节数组。

如果只是装字节数组那程序非常好写，直接ReadBytes(attribute_length)即可，但这个info虽然是字节数组，但其也是一个数据类型，官方对info定义了23种类型，info结构体的解析非常之复杂，我简单的说一下我解析info中的Code和StackMapTable的思路：

### Code
官方对Code结构体定义如下：

```c
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```

前面的字段均为u2或u4，之后会有一个异常表的结构体(ExceptionTable)

我定义异常表的Read方法，然后实现，在需要读取的时候直接调用即可，以下可供参考：

```go
type Code_attribute struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	MaxStrack            uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionTable
	AttributesCount      uint16
	Attributes           Attributes
}

type ExceptionTable struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (e *ExceptionTable) readExceptionTable(reader class_file_commons.Reader) ExceptionTable {
	e.StartPc = reader.ReadUint16()
	e.EndPc = reader.ReadUint16()
	e.HandlerPc = reader.ReadUint16()
	e.CatchType = reader.ReadUint16()
	return *e
}

func (c *Code_attribute) ReadAttrInfo(reader class_file_commons.Reader) AttrInfo {
	c.MaxStrack = reader.ReadUint16()
	c.MaxLocals = reader.ReadUint16()
	c.CodeLength = reader.ReadUint32()
	c.Code = reader.ReadBytes(c.CodeLength)
	c.ExceptionTableLength = reader.ReadUint16()
	size := c.ExceptionTableLength
	for i := 0; i < int(size); i++ {
		c.ExceptionTable = append(c.ExceptionTable, new(ExceptionTable).readExceptionTable(reader))
	}
	ar := AttributesStructReader{reader}
	a_count, a_infos := ar.ReadAttributeInfos()
	c.AttributesCount = a_count
	c.Attributes = *a_infos
	return c
}
```

### StackMapTable
关于StackMapTable的结构类型比较复杂，代码量也是其他结构体的3倍左右，因为其子类型太多了

官方对StackMapTable定义如下：

```c
StackMapTable_attribute {
    u2              attribute_name_index;
    u4              attribute_length;
    u2              number_of_entries;
    stack_map_frame entries[number_of_entries];
}
```

其中stack_map_frame定义如下：

```c
union stack_map_frame {
    same_frame; // 0-63
    same_locals_1_stack_item_frame; // 64-127
    same_locals_1_stack_item_frame_extended; // 247
    chop_frame; // 248-250
    same_frame_extended; // 251
    append_frame; // 252-254
    full_frame; // 255
}
```

可以看到是一个联合体，也就是说这个stack_map_frame之后还需要再次定义结构体，比如same_frame、full_frame等等，这些官网皆有说明，摘取一个简单的：

```c
same_frame {
    u1 frame_type = SAME; /* 0-63 */
}
```

这个意思表示当偏移值为0-63时，stack_map_frame为same_frame类型。这个偏移值由第一个字节表示。

所以思路就是先读一个字节，然后判断这个字节的范围确定对应的结构体类型。

但其实之后还有一个结构体：VerificationTypeInfo，定义如下：

```c
union verification_type_info {
    Top_variable_info; // 0
    Integer_variable_info; // 1
    Float_variable_info; // 2
    Long_variable_info; // 4
    Double_variable_info; // 3
    Null_variable_info; // 5
    UninitializedThis_variable_info; // 6
    Object_variable_info; // 7
    Uninitialized_variable_info; // 8
}
```

可以看到还是联合体。。。所以其实这一个StackMapTable就需要定义26种数据类型，这对于go来说着实不太友好，因为go不面向对象，但我经过3天的思考和对代码的review，将代码设计成可控的形式了，具体实现可以参考(jvm/class/attribute/attr_info_StackMapTable.go)。

用这个方法我可以在我自己的维度中解决go不面向对象的问题。


## 3.8 解析ClassFile结构
综上内容，可以实现对一个类结构体的解析

由于参数表数目众多且多数为冗余，我并没有全部实现，未实现我直接读取字节码便于后续处理。

代码参考(jvm/class/class_fil.go)

使用如下参数：
```
.\jvm-0.1.0.exe -cp T:\\jvm-test -p --no-trunc java.lang.String aaa bbb
```

运行后发现输出了字节码的详细格式

# 参考
- [StackMapTable结构体的解析](https://blog.csdn.net/r77683962/article/details/78877228?utm_medium=distribute.pc_relevant_bbs_down.none-task-blog-baidujs-1.nonecase&depth_1-utm_source=distribute.pc_relevant_bbs_down.none-task-blog-baidujs-1.nonecase)