package pojo

type ZSetItem struct {
	val   string
	score string
}

func NewZSetItem() *ZSetItem { // 返回结构体ZSetItem实例的指针
	item := new(ZSetItem)
	return item
}

type ZSetRes struct {
	set []ZSetItem
}

func NewZSetRes() *ZSetRes { // 返回结构体ZSetRes实例的指针
	res := new(ZSetRes)
	return res
}
