package constant_pool

type TagInfo struct {
	Tag uint8
}

func (t *TagInfo) GetTag() uint8 {
	return t.Tag
}

func (t *TagInfo) GetTagName() string {
	return CONSTANT_POOL_TAG_NAME_MAP[t.GetTag()]
}
