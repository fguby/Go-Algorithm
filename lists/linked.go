package lists


type Linked interface {
	//判断是否为空
	IsEmpty() bool
	//清空线性表
	ClearList()
	//返回链表中的第i个元素
	GetElem(i int) (interface{}, error)
	//从头部添加元素
	Insert(e interface{})
	//从尾部添加元素
	Append(e interface{})
	//从序号处添加元素
	InsertNode(i int, e interface{}) error
	//删除第i个位置处元素
	Remove(i int) error
	//返回size
	Length() int
}