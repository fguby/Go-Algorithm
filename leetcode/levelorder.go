package leetcode

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var nums []*TreeNode

	nums = append(nums, root)

	var results [][]int

	for len(nums) > 0 {

		var arr []int

		var news []*TreeNode

		for _, v := range nums {
			arr = append(arr, v.Val)
			if v.Left != nil {
				news = append(news, v.Left)
			}
			if v.Right != nil {
				news = append(news, v.Right)
			}
		}

		results = append(results, arr)

		nums = news
	}

	return results
}
