// package tree

// //给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

// /**
// 输入: 3
// 输出:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// 解释:
// 以上的输出对应以下 5 种不同结构的二叉搜索树：

//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3
// **/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func generateTrees(n int) []*TreeNode {

// 	trees := make([]*TreeNode, 0, 2*n)

// 	for n > 0 {
// 		t := &TreeNode{
// 			Val: n,
// 			Left: nil,
// 			Right: nil,
// 		}

// 		i := n
// 		for i > 0 {

// 		}
// 		trees = append(trees, insert(t,))

// 		n --
// 	}

// 	return 
// }

// func insert(t *TreeNode, n int) []*TreeNode {

// 	if t == nil {
// 		return []*TreeNode{
// 			&TreeNode{
// 				Val:   n,
// 				Left:  nil,
// 				Right: nil,
// 			}}
// 	}

// 	var left []*TreeNode
// 	var right []*TreeNode

// 	if n < t.Val {
// 		left = append(left, insert(t.Left, n)...)
// 	} else {
// 		right = append(right, insert(t.Right, n)...)
// 	}

// 	for _, v := range left {
// 		//右旋转
// 		if v.Left != nil && v.Right == nil {

// 			left = append(left, leftRotate(v))

// 		} else if v.Right != nil && v.Left == nil {

// 			left = append(left, rightRotate(v))
// 		}
// 	}

// 	for _, v := range right {
// 		if v.Left != nil && v.Right == nil {

// 			right = append(right, leftRotate(v))

// 		} else if v.Right != nil && v.Left == nil {

// 			right = append(right, rightRotate(v))
// 		}
// 	}

// 	result := make([]*TreeNode, 0, len(left)+len(right))

// 	if len(left) == 0 {

// 		for _, v := range right {
// 			newTree := t
// 			newTree.Right = v
// 			result = append(result, newTree)
// 		}

// 	} else if len(right) == 0 {

// 		for _, v := range left {
// 			newTree := t
// 			newTree.Left = v
// 			result = append(result, newTree)
// 		}

// 	} else {

// 		for _, l := range left {
// 			for _, r := range right {
// 				newTree := t
// 				newTree.Left = l
// 				newTree.Right = r
// 				result = append(result, newTree)
// 			}
// 		}
// 	}

// 	return result
// }

// func leftRotate(v *TreeNode) *TreeNode {

// 	m := v.Left
// 	v.Left = m.Right
// 	m.Right = v

// 	return m
// }

// func rightRotate(v *TreeNode) *TreeNode {

// 	m := v.Right
// 	v.Right = m.Left
// 	m.Left = v

// 	return m
// }
