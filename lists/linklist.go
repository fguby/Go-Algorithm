package lists

import (
	"fmt"
)

type IndexValidError struct {
	index int
}

func (iErr IndexValidError) Error() string {
	return fmt.Sprintf("Index is illegal:%d", iErr.index)
}

type LinkList struct {
	headNode *Node //头节点
}

type Node struct {
	Data interface{}
	Next *Node
}

//设立一个虚拟头结点
func NewLinkedList() *LinkList {
	return &LinkList{headNode: &Node{
		Data: nil,
		Next: nil,
	}}
}

func assertLinkListImplementation() {
	var _ Linked = (*LinkList)(nil)
}

func (l *LinkList) IsEmpty() bool {
	if l.headNode.Next == nil {
		return true
	}
	return false
}

func (l *LinkList) Length() int {
	cur := l.headNode
	count := 0

	for cur.Next != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (l *LinkList) Insert(e interface{}) {
	node := &Node{
		Data: e,
	}
	node.Next = l.headNode.Next
	l.headNode.Next = node
}

func (l *LinkList) Append(e interface{}) {
	node := &Node{
		Data: e,
	}

	if l.IsEmpty() {
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (l *LinkList) InsertNode(i int, e interface{}) error {
	if i <= 0 || i > l.Length() {
		return IndexValidError{index: i}
	}

	cur := l.headNode

	for o := 0; o < i; o++ {
		cur = cur.Next
	}

	node := &Node{
		Data: e,
	}
	node.Next = cur.Next
	cur.Next = node
	return nil
}

func (l *LinkList) Remove(i int) error {

	if i <= 0 || i > l.Length() {
		return IndexValidError{index: i}
	}

	cur := l.headNode
	for o := 0; o < i-1; o++ {
		cur = cur.Next
	}
	if i == l.Length() {
		cur.Next = nil
	} else {
		cur.Next = cur.Next.Next
	}
	return nil
}

func (l *LinkList) GetElem(i int) (interface{}, error) {
	if i < 0 || i > l.Length() {
		return nil, IndexValidError{index: i}
	}

	cur := l.headNode
	for o := 0; o < i; o++ {
		cur = cur.Next
	}
	return cur.Data, nil
}

func (l *LinkList) ClearList() {
	l.headNode.Data = nil
	l.headNode.Next = nil
}

func (l *LinkList) GetHeader() *Node {
	return l.headNode
}

//使用递归算法来移除元素
func RemoveValue(node *Node, e interface{}) *Node {

	if node.Next == nil {
		return nil
	}

	node.Next = RemoveValue(node.Next, e)

	if node.Data == e {
		return node.Next
	}
	return node
}
