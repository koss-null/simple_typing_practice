package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	rightIdx := 0
	for i := 0; i <= rightIdx; i++ {
		if rightIdx >= len(nums) {
			return true
		}
		rightIdx += nums[i]
	}
	return false
}
