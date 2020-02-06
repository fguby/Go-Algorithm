package trees

import (
	"fmt"
)

/**
线段树
常用来解决区间问题
*/

type merge = func(e1, e2 interface{}) interface{}

type SegmentTree struct {
	m    merge
	arr  []interface{}
	data []interface{}
}

//将一个数组排列成线段树
func NewSegmentTree(nums []interface{}, fn merge) *SegmentTree {
	if len(nums) == 0 {
		return nil
	}
	newNums := make([]interface{}, 4*len(nums))
	tree := &SegmentTree{
		m:    fn,
		data: nums,
		arr:  newNums,
	}
	//查找最后一个非叶子节点
	tree.buildSegmentTree(0, 0, len(nums)-1)
	return tree
}

//在索引treeIndex处创建l - r的区间元素
func (tree *SegmentTree) buildSegmentTree(treeIndex int, l int, r int) {
	if l == r {
		tree.arr[treeIndex] = tree.data[l]
		return
	}
	leftChildIndex, _ := tree.GetLeftChild(treeIndex)
	rightChildIndex, _ := tree.GetRightChild(treeIndex)

	mid := l + (r-l)/2
	tree.buildSegmentTree(leftChildIndex, l, mid)
	tree.buildSegmentTree(rightChildIndex, mid+1, r)

	tree.arr[treeIndex] = tree.m(tree.arr[leftChildIndex], tree.arr[rightChildIndex])
}

//获取k索引节点处的左孩子
func (tree *SegmentTree) GetLeftChild(k int) (int, error) {
	leftChildIndex := 2*k + 1
	if leftChildIndex >= len(tree.arr) {
		return -1, fmt.Errorf("invalid index:")
	}
	return leftChildIndex, nil
}

//获取k索引节点处的右孩子
func (tree *SegmentTree) GetRightChild(k int) (int, error) {
	rightChildIndex := 2*k + 2
	if rightChildIndex >= len(tree.arr) {
		return -1, fmt.Errorf("invalid index:")
	}
	return rightChildIndex, nil
}

func (tree *SegmentTree) GetParent(k int) (int, error) {
	if k <= 0 {
		return -1, fmt.Errorf("invalid index:")
	}
	return (k - 1) / 2, nil
}

func (tree *SegmentTree) GetRootNode() interface{} {
	return tree.arr[0]
}

func (tree *SegmentTree) Update(k int, v interface{}) {
	tree.data[k] = v
	tree.update(0, 0, len(tree.data)-1, k)
}

func (tree *SegmentTree) update(treeIndex, l, r, k int) {

	if l == r && l == k {
		tree.arr[treeIndex] = tree.data[k]
		return
	}

	midle := l + (r-l)/2

	leftChildIndex, _ := tree.GetLeftChild(treeIndex)
	rightChildeIndex, _ := tree.GetRightChild(treeIndex)

	if k <= midle {
		tree.update(leftChildIndex, l, midle, k)
	} else {
		tree.update(rightChildeIndex, midle+1, r, k)
	}

	tree.arr[treeIndex] = tree.m(
		tree.arr[leftChildIndex],
		tree.arr[rightChildeIndex],
	)
}

func (tree *SegmentTree) Query(queryL, queryR int) interface{} {
	return tree.query(0, 0, len(tree.data)-1, queryL, queryR)
}

/**
在以treeIndex为根的区间树【l...r】范围里
搜索queryL到queryR的值
*/
func (tree *SegmentTree) query(treeIndex, l, r, queryL, queryR int) interface{} {
	//queryL必须小于等于queryR
	if queryL > queryR {
		return nil
	}

	if queryL == queryR {
		return tree.data[queryL]
	}
	//区间完全吻合的情况
	if l == queryL && r == queryR {
		return tree.arr[treeIndex]
	}
	//查看当前区间的中间值
	midle := l + (r-l)/2
	//获取当前节点的右孩子
	rightIndex, _ := tree.GetRightChild(treeIndex)
	//获取当前节点的左孩子
	leftIndex, _ := tree.GetLeftChild(treeIndex)
	if queryL >= (midle + 1) {
		return tree.query(rightIndex, (midle + 1), r, queryL, queryR)
	} else if queryR <= midle {
		return tree.query(leftIndex, l, midle, queryL, queryR)
	} else {
		//区间值横跨左右两个孩子
		return tree.m(
			tree.query(leftIndex, l, midle, queryL, midle),
			tree.query(rightIndex, (midle+1), r, (midle+1), queryR),
		)
	}
}
