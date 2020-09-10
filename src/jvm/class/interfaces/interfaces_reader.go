package interfaces

import "jvm/class/class_file_commons"

type InterfacesStructReader struct {
	class_file_commons.Reader
}

func (ir *InterfacesStructReader) ReadInterfacesInfos() (int, []uint16) {
	ret := *new([]uint16)
	size := ir.ReadUint16()
	for i := 0; i < int(size); i++ {
		ret = append(ret, ir.ReadUint16())
	}
	return int(size), ret
}
