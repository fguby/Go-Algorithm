package leetcode

func canJump(nums []int) bool {

	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	recent := len(nums) - 1

	for i := len(nums) - 1; i >= 0; {

		if nums[i] >= recent-i {
			recent = i
		}

		i--
	}

	return recent == 0
}
