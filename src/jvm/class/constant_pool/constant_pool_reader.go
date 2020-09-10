package constant_pool

import (
	"jvm/class/class_file_commons"
)

type ConstantPoolStructReader struct {
	class_file_commons.Reader
}

/**
e.g	解析常量池需要注意如下：
	1. 常量池实际大小比constant_pool_size要小1
	1. Double和Long各占2个位置，也就是说解析Double和Long时，它们的下一个位置为nil
*/
func (c *ConstantPoolStructReader) ReadConstantPoolInfos() *ConstantPool {
	ret := new(ConstantPool)
	size := c.ReadUint16()
	ret.Size = size
	for i := 0; i < int(size)-1; i++ {
		// 遍历读取常量池中元素，常量池下标从1开始一直到size-1结束
		tag := c.ReadUint8()
		handle := newConstantInfo(tag)
		info := handle.ReadInfo(c)
		ret.ConstantItemInfos = append(ret.ConstantItemInfos, info)
		if handle.GetTag() == CONSTANT_Long || handle.GetTag() == CONSTANT_Double {
			i++
			ret.ConstantItemInfos = append(ret.ConstantItemInfos, nil)
		}
	}
	return ret
}
