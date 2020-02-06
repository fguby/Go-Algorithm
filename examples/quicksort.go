package examples

/**
快速排序算法
*/
func QuickSort(nums []int) {

	if len(nums) < 2 {
		return
	}

	cur, l := 0, len(nums)-1

	for i := 1; i <= l; {
		if nums[i] <= nums[cur] {
			nums[cur], nums[i] = nums[i], nums[cur]
			cur = i
			i++
		} else {
			nums[i], nums[l] = nums[l], nums[i]
			l--
		}
	}

	QuickSort(nums[:cur])

	QuickSort(nums[cur+1:])

}
