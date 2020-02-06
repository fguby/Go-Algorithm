package dynamic_programming

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

/**
	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。

	输入: "cbbd"
	输出: "bb"
**/

//解题思路
/**
	//暴力破解法
	for
		for
		//拼接字符串的时候判断是否是回文
	时间复杂度为n^3级别


	//动态规划解法
	s[i, j]如果是回问子串的话，那么s[i+1, j-1]也必定是回文子串

**/
//动态规划解法
func LongestPalindrome(s string) string {

	isPalindrome := make([][]bool, len(s))

	maxLength := 1

	maxStr := string(s[0])

	for k := range isPalindrome {
		isPalindrome[k] = append(isPalindrome[k], make([]bool, len(s))...)
	}

	//长度为1的必然是回文子串，从长度为2开始循环
	for l := 2; l <= len(s); l++ {

		for i := 0; i < len(s)-1; i++ {

			//遍历到底，退出
			if i == len(s)-(l-1) {
				break
			}
			//判断首尾字母是否相同
			if s[i] == s[i+l-1] {

				start := i + 1
				end := i + l - 2

				//判断边界
				if start >= end {
					isPalindrome[i][i+l-1] = true
				} else {
					if isPalindrome[start][end] {
						isPalindrome[i][i+l-1] = true
					} else {
						isPalindrome[i][i+l-1] = false
					}
				}
				//是回文子串
				if isPalindrome[i][i+l-1] {
					length := (i + l - 1) - i + 1
					if length > maxLength {
						maxLength = length
						maxStr = s[i : i+l]
					}
				}
			} else {
				isPalindrome[i][i+l-1] = false
			}
		}
	}

	return maxStr
}

//Manacher算法
//TODO
