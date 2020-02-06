package leetcode

/**
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
