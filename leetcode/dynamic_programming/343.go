package dynamic_programming

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。
//返回你可以获得的最大乘积。

/**
示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。

示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。

说明: 你可以假设 n 不小于 2 且不大于 58。
**/

//递归解法
func IntegerBreak(n int) int {

	if n == 1 {
		return n
	}

	result := -1

	//最大数
	//
	for i := 1; i < n; i++ {

		result = maxThreeNumber(result, i*(n-i), i*IntegerBreak(n-i))
	}

	return result
}

//动态规划解法
func IntegerBreak1(n int) int {

	nums := make([]int, n)

	nums[0] = 1
	nums[1] = 2

	result := nums[0]

	for i := 3; i <= n; i++ {

		for y := 1; y < i; y++ {
			result = maxThreeNumber(result, y*(i-y), y*nums[i-y-1])
		}
		nums[i-1] = result
	}

	return result
}

func maxThreeNumber(n1, n2, n3 int) int {
	if n1 > n2 {
		if n1 > n3 {
			return n1
		} else {
			return n3
		}
	} else {
		if n2 > n3 {
			return n2
		} else {
			return n3
		}
	}
}
