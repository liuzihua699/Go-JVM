package constant_pool

import (
	"jvm/class/class_file_commons"
)

type ConstantPoolStructReader struct {
	class_file_commons.Reader
}

/**
e.g	解析常量池需要注意如下：
	1. 常量池下标从1开始，官方规定0为无效索引
	2. 常量池实际大小比constant_pool_size要小1
	3. Double和Long各占2个位置，也就是说解析Double和Long时，它们的下一个位置为nil
	4. 需要优先解析URF8常量并放置为 index:name 的形式
*/
func (c *ConstantPoolStructReader) ReadConstantPoolInfos() (int, *ConstantPool) {
	ret := new(ConstantPool)
	size := c.ReadUint16()
	ret.Utf8Map = make(map[uint16]string)

	// 遍历读取常量池中元素，常量池下标从1开始一直到size-1结束
	for i := 0; i < int(size)-1; i++ {
		tag := c.ReadUint8()
		info := newConstantInfo(tag).ReadInfo(c)
		ret.ConstantItemInfos = append(ret.ConstantItemInfos, info)
		if info.GetTag() == CONSTANT_Long || info.GetTag() == CONSTANT_Double {
			i++
			ret.ConstantItemInfos = append(ret.ConstantItemInfos, nil)
		}
		if info.GetTag() == CONSTANT_Utf8 {
			bytes := info.(*ConstantUtf8Info).Bytes
			ret.Utf8Map[uint16(i+1)] = string(bytes)
		}
	}
	setFileConstantPool(ret)
	return int(size), ret
}
