package trees

import (
	"GoAlgorithm/utils"
	"math"
)

type AvlRootTree struct {
	root        *avltree
	compare     utils.Comparator
	accumulator utils.Accumulator
}

type avltree struct {
	node   mapNode
	left   *avltree
	right  *avltree
	height int
}

type mapNode struct {
	k     interface{}
	value interface{}
}

// func assertAvlTreeImplementation() {
// 	var _ tree = (*AvlRootTree)(nil)
// }

func NewAvlRootTree(fn utils.Comparator) *AvlRootTree {
	return &AvlRootTree{
		root:    nil,
		compare: fn,
	}
}

func NewAvlTreeWithIntComparator() *AvlRootTree {
	return &AvlRootTree{
		root:        nil,
		compare:     utils.IntComparator,
		accumulator: utils.IntAccumulator,
	}
}

func NewAvlTreeWithStringComparator() *AvlRootTree {
	return &AvlRootTree{
		root:        nil,
		compare:     utils.StringComparator,
		accumulator: utils.IntAccumulator,
	}
}

//判断是否为空
func (avl *AvlRootTree) IsEmpty() bool {
	return avl.root == nil
}

//取得根节点
func (avl *AvlRootTree) GetRootNode() interface{} {
	return avl.root
}

//
func (avl *AvlRootTree) Insert(k, v interface{}) {
	if avl.root == nil {
		avl.root = &avltree{
			node: mapNode{
				k:     k,
				value: v,
			},
			left:   nil,
			right:  nil,
			height: 1,
		}
	} else {
		avl.root = avl.root.Add(k, v, avl.compare, avl.accumulator)
	}
}

func (tree *avltree) Add(k, v interface{},
	fn1 utils.Comparator, fn2 utils.Accumulator) *avltree {

	//递归条件终止
	if tree == nil {
		return &avltree{
			node: mapNode{
				k:     k,
				value: v,
			},
			left:   nil,
			right:  nil,
			height: 1,
		}
	}

	switch value, _ := fn1(tree.node.k, k); value {
	case 1:
		tree.left = tree.left.Add(k, v, fn1, fn2)
	case -1:
		tree.right = tree.right.Add(k, v, fn1, fn2)
	default:
		tree.node.value, _ = fn2(tree.node.value, v)
		return tree
	}

	balance := tree.GetBalanceFactor()

	//平衡因子超过1，需要进行旋转操作
	if math.Abs(float64(balance)) > 1 {

		//新增节点类型——RR
		if balance < 0 && tree.right.GetBalanceFactor() <= 0 {
			tree = tree.leftRotate()
		}

		//新增节点类型——RL
		if balance < 0 && tree.right.GetBalanceFactor() > 0 {
			tree.right = tree.right.rightRotate()
			tree = tree.leftRotate()
		}

		//新增节点类型——LL
		if balance > 0 && tree.left.GetBalanceFactor() >= 0 {
			tree = tree.rightRotate()
		}

		//新增节点类型——LR
		if balance > 0 && tree.left.GetBalanceFactor() < 0 {
			tree.left = tree.left.leftRotate()
			tree = tree.rightRotate()
		}

	}

	//
	tree.height = tree.GetMaxHeight() + 1

	return tree
}

//左旋转
func (tree *avltree) leftRotate() *avltree {

	p := tree.right

	tree.right = p.left
	p.left = tree
	//更新高度值
	p.left.height = p.left.GetMaxHeight() + 1
	p.height = p.GetHeight() + 1
	return p
}

//右旋转
func (tree *avltree) rightRotate() *avltree {
	p := tree.left
	tree.left = p.right
	p.right = tree

	//更新高度值
	p.right.height = p.right.GetMaxHeight() + 1
	p.height = p.GetMaxHeight() + 1
	return p
}

//计算平衡因子
func (tree *avltree) GetBalanceFactor() int {
	if tree == nil {
		return 0
	}
	return tree.left.GetHeight() - tree.right.GetHeight()
}

//获取高度值
func (tree *avltree) GetHeight() int {
	if tree == nil {
		return 0
	}
	return tree.height
}

func (tree *avltree) GetMaxHeight() int {
	if tree.left.GetHeight() > tree.right.GetHeight() {
		return tree.left.GetHeight()
	}
	return tree.right.GetHeight()
}

//判断是否是二分树
func (avl *AvlRootTree) IsBinaryTree() bool {
	nums := make([]interface{}, 0, avl.root.height*20)
	//中序遍历
	avl.root.Inorder(nums)

	for k, _ := range nums {
		if value, _ := avl.compare(nums[k], nums[k+1]); value > 0 {
			return false
		}
	}
	return true
}

func (avl *AvlRootTree) Remove(k interface{}) {
	avl.root = avl.root.remove(k, avl.compare)
}

func (tree *avltree) remove(k interface{}, fn utils.Comparator) *avltree {

	if tree == nil {
		return nil
	}

	value, _ := fn(tree.node.k, k)

	if value > 0 {
		tree.left = tree.left.remove(k, fn)
	} else if value < 0 {
		tree.right = tree.right.remove(k, fn)
	} else {

		p := tree.right

		//找到要删除的节点
		if tree.left == nil {
			tree.right = nil
		} else if tree.right == nil {
			p = tree.left
			tree.left = nil
		} else {
			//右子树中的最小值
			node := tree.right.getMinimum()
			tree.node = node.node
			tree.right = tree.right.remove(node.node.k, fn)
			p = tree
		}
		return p
	}

	//平衡检测
	balance := tree.GetBalanceFactor()

	//平衡因子超过1，需要进行旋转操作
	if math.Abs(float64(balance)) > 1 {

		//新增节点类型——RR
		if balance < 0 && tree.right.GetBalanceFactor() <= 0 {
			tree = tree.leftRotate()
		}

		//新增节点类型——RL
		if balance < 0 && tree.right.GetBalanceFactor() > 0 {
			tree.right = tree.right.rightRotate()
			tree = tree.leftRotate()
		}

		//新增节点类型——LL
		if balance > 0 && tree.left.GetBalanceFactor() >= 0 {
			tree = tree.rightRotate()
		}

		//新增节点类型——LR
		if balance > 0 && tree.left.GetBalanceFactor() < 0 {
			tree.left = tree.left.leftRotate()
			tree = tree.rightRotate()
		}

	}

	//
	tree.height = tree.GetMaxHeight() + 1

	return tree
}

func (tree *avltree) getMinimum() *avltree {

	if tree.left == nil {
		return tree
	}

	return tree.left.getMinimum()
}

func (tree *avltree) Inorder(nums []interface{}) {
	if tree == nil {
		return
	}

	tree.left.Inorder(nums)

	nums = append(nums, tree.node.k)

	tree.right.Inorder(nums)
}
