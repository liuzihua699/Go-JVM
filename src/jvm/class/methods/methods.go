package fields

import (
	"jvm/class/attribute"
)

/**
method_info {
	u2             access_flags;
	u2             name_index;
	u2             descriptor_index;
	u2             attributes_count;
	attribute_info attributes[attributes_count];
}
*/
type Methods struct {
	MemberInfos []attribute.MemberInfo
}
