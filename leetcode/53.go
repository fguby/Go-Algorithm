package leetcode

/**
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
**/

/**
	按层遍历
	最外层-次外层....最里层
	设定边界点
**/
func SpiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	results := make([]int, 0, len(matrix)*len(matrix[0]))

	r1, r2 := 0, len(matrix)-1
	c1, c2 := 0, len(matrix[0])-1

	for (r1 <= r2) && (c1 <= c2) {
		for c := c1; c <= c2; c++ {
			results = append(results, matrix[r1][c])
		}

		for r := r1 + 1; r <= r2; r++ {
			results = append(results, matrix[r][c2])
		}
		if c1 < c2 && r1 < r2 {
			for c := c2 - 1; c > c1; c-- {
				results = append(results, matrix[r2][c])
			}

			for r := r2; r > r1; r-- {
				results = append(results, matrix[r][c1])
			}
		}

		r1++
		r2--
		c1++
		c2--
	}

	return results
}
