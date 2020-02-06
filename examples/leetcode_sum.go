package examples

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
找出所有满足条件且不重复的三元组。
例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func ThreeSum(nums []int) [][]int {

	//小于三个元素返回
	if len(nums) < 3 {
		return nil
	}

	var arr [][]int

	//排序
	QuickSort(nums)

	//整个数组不成立情况过滤
	if nums[0] > 0 ||
		nums[len(nums)-1] < 0 ||
		nums[0]+nums[1]+nums[2] > 0 ||
		nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < 0 {
		return nil
	}

	for i := 0; i < len(nums)-1; i++ {
		r, l := i+1, len(nums)-1
		//设置一个中间数
		middleNum := l
		if l != (r + 1) {
			middleNum = ((l - r) / 2) + r
		}

		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//当前元素大于0，则之后的所有组合必都大于0
		if nums[i] > 0 {
			break
		}

		for r < middleNum && middleNum <= l {

			//优化不可能的情况
			if r > i+1 && nums[r] == nums[r-1] {
				r++
				continue
			}

			if (nums[i] + nums[r] + nums[middleNum]) > 0 {
				//middle走到最右
				if middleNum == l {
					r++
				}
				middleNum--
			} else if (nums[i] + nums[r] + nums[middleNum]) < 0 {
				if middleNum == l {
					r++
					continue
				}
				middleNum++
			} else {
				arr = append(arr, []int{nums[i], nums[r], nums[middleNum]})
				r++
				l--
				if l == (r + 1) {
					middleNum = l
				} else {
					middleNum = ((l - r) / 2) + r
				}
			}
		}
	}

	return arr
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	QuickSort(nums) // 快排 O(n log n), 本地修改

	ret := make([][]int, 0, 0)
	for i := 0; i < len(nums)-1; i++ {
		l, r := i+1, len(nums)-1

		if nums[i] > 0 || nums[i]+nums[l] > 0 {
			break
		} // 优化掉不可能的情况

		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // i 去重

		for l < r {
			if l > i+1 && nums[l] == nums[l-1] { // l 去重
				l++
				continue
			}
			if r < len(nums)-2 && nums[r] == nums[r+1] { // r 去重
				r--
				continue
			}

			if nums[i]+nums[l]+nums[r] > 0 { //
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}

	return ret
}
