package leetcode

import (
	"bytes"
	"strings"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func LengthOfLongestSubstring1(s string) int {

	nums := [256]int{}

	max := 0

	cur := 0

	for k, v := range s {

		if nums[v] > 0 && nums[v] > cur {
			cur = nums[v]
		}

		nums[v] = k + 1

		max = Max(max, k-cur+1)
	}

	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LengthOfLongestSubstring(s string) int {

	var str string

	var bt bytes.Buffer

	m := &max{0}

	for _, v := range s {

		value := strings.IndexRune(str, v)

		if value != -1 {

			if value == len(str)-1 {
				m.Add(len(str))
				s = ""
				bt.Reset()
			} else {
				str = str[value+1:]
				b := bt.Bytes()[value+1:]
				bt.Reset()
				bt.Write(b)
				m.Add(value + 1)
			}
		}

		bt.WriteRune(v)
		str = bt.String()
		m.Add(len(str))
	}
	//m.Add(len(str))
	return m.val
}

type max struct {
	val int
}

func (m *max) Add(value int) {
	if m.val < value {
		m.val = value
	}
}

/**
func lengthOfLongestSubstring(s string) int {
    flag := [256]int{}
    beg := 0
    max := 0
    for i := 0; i < len(s); i++ {
        if flag[s[i]] > 0 && flag[s[i]] > beg {
            beg = flag[s[i]]
        }
        flag[s[i]] = i+1
        max = maxnum(max, i - beg + 1)
    }
    return max
}

func maxnum(a, b int) int {
    if a > b {
        return a
    }
    return b
}
*/
