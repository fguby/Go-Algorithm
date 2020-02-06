package trees

import (
	"GoAlgorithm/queue"
	"GoAlgorithm/stack"
	"GoAlgorithm/utils"
	"fmt"
)

type eNode = interface{}

//BinarySearchTree
type Bst struct {
	root       *node
	comparator utils.Comparator
	size       int
}

//Tree
type node struct {
	val   eNode
	left  *node
	right *node
	depth int
}

func assertBstImplementation() {
	var _ tree = (*Bst)(nil)
}

//
func NewWith(c utils.Comparator) *Bst {
	return &Bst{
		root:       nil,
		comparator: c,
	}
}

func NewWithIntComparator() *Bst {
	return &Bst{
		root:       nil,
		comparator: utils.IntComparator,
	}
}

func (b *Bst) IsEmpty() bool {
	//判断根节点是否为空即可
	if b.root == nil {
		return true
	}
	return false
}

func (b *Bst) GetSize() int {
	return b.size
}

func (b *Bst) Insert(e interface{}) {
	if b.root == nil {
		b.root = &node{
			val:   e,
			left:  nil,
			right: nil,
		}
		b.size = 1
	} else {
		b.insertNode(b.root, e, 0)
	}
}

func (b *Bst) insertNode(n *node, e interface{}, depth int) *node {

	if n == nil {
		n = &node{
			val:   e,
			left:  nil,
			right: nil,
			depth: depth + 1,
		}
		b.size++
		return n
	}

	v, _ := b.comparator(n.val, e)

	if v > 0 {
		n.left = b.insertNode(n.left, e, n.depth)
	} else if v < 0 {
		n.right = b.insertNode(n.right, e, n.depth)
	}

	return n
}

func (b *Bst) GetRootNode() eNode {
	return b.root.val
}

func (b *Bst) Containers(e interface{}) eNode {

	if b.root == nil {
		return nil
	}

	v := b.searchNode(b.root, e)
	return v.val
}

func (b *Bst) searchNode(n *node, e interface{}) *node {

	if n == nil {
		return nil
	}

	v, _ := b.comparator(n.val, e)

	if v > 0 {
		n = b.searchNode(n.left, e)
	} else if v < 0 {
		n = b.searchNode(n.right, e)
	}
	return n
}

//前序遍历
func (b *Bst) PreOrder() {
	preOrder(b.root)
}

//非递归方式前序遍历
func (b *Bst) PreOrderNR() {

	s := stack.NewArrayStack(15)
	s.Push(b.root)

	for s.GetSize() > 0 {
		//
		n, _ := s.Pop()

		node := n.(*node)

		fmt.Println(node.val)
		fmt.Println(node.depth)
		fmt.Println("----")

		if node.right != nil {
			s.Push(node.right)
		}

		if node.left != nil {
			s.Push(node.left)
		}
	}
}

//层序遍历
func (b *Bst) LevelOrder() {
	q := queue.NewArrayQueue(15)
	q.Enqueue(b.root)

	for !q.IsEmpty() {

		value, _ := q.Dequeue()

		n := value.(*node)

		fmt.Println(n.val)

		if n.left != nil {
			q.Enqueue(n.left)
		}

		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
}

//中序遍历
func (b *Bst) InOrder() {
	inOrder(b.root)
}

//后序遍历
func (b *Bst) LastOrder() {
	lastOrder(b.root)
}

//删除最小元素
func (b *Bst) DeleteMin() {

	b.size--

	//最小元素为根节点时
	if b.root.left == nil {
		b.root = b.root.right
		return
	}
	deleteMin(b.root)
}

//删除最大元素
func (b *Bst) DeleteMax() {

	b.size--

	if b.root.right == nil {
		b.root = b.root.left
		return
	}
	deleteMax(b.root)
}

//删除任意节点
func (b *Bst) DeleteNode(e interface{}) {
	b.deleteNode(b.root, e, b.comparator)
}

func preOrder(n *node) {

	if n == nil {
		return
	}

	fmt.Println(n.val)

	preOrder(n.left)

	preOrder(n.right)
}

func inOrder(n *node) {
	if n == nil {
		return
	}
	inOrder(n.left)

	fmt.Println(n.val)

	inOrder(n.right)
}

func lastOrder(n *node) {
	if n == nil {
		return
	}

	lastOrder(n.left)

	lastOrder(n.right)

	fmt.Println(n.val)
}

func deleteMin(n *node) *node {
	if n.left != nil {
		n.left = deleteMin(n.left)
	} else {
		n = n.right
	}
	return n
}

func deleteMax(n *node) {
	if n.right.right != nil {
		deleteMax(n.right)
	} else {
		n.right = n.right.left
	}
}

func (b *Bst) deleteNode(n *node, e interface{}, c utils.Comparator) *node {

	if n == nil {
		return nil
	}

	value, _ := c(n.val, e)

	if value > 0 {
		n.left = b.deleteNode(n.left, e, c)
	} else if value < 0 {
		n.right = b.deleteNode(n.right, e, c)
	} else {
		//
		if n.left == nil && n.right != nil {
			n = n.right
		} else if n.right == nil && n.left != nil {
			n = n.left
		} else if n.right == nil && n.left == nil {
			n = nil
		} else {
			/**
			左右子树都不为空
			寻找右子树里的最小值
			*/
			rightN := n.right
			for rightN.left != nil {
				rightN = rightN.left
			}
			n.val = rightN.val
			n.right = deleteMin(n.right)
		}
		b.size--
	}
	return n
}
