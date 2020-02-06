package heap

/**
	heap 堆

1.完全二叉树（每一行节点都从左到右排满。）
2.任一节点的父节点都必然大于等于它。
3.可以用数组表示。

父亲节点： （k - 1）/ 2
左孩子节点： 2*k + 1
右孩子节点:  2*k + 2

*/
type heap interface {
	GetParent(k int) (int, error)
	GetSize() int
	IsEmpty() bool
	GetLeftChild(k int) (int, error)
	GetRightChild(k int) (int, error)
	SiftUp(k int)
	SiftDown(k int)
	Replace(k int, e interface{})
	Delete(k int)
}
