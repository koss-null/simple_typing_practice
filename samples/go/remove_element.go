func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	red := 0
	unique_cnt := 0
	for red != len(nums) {
		for nums[unique_cnt] != val && unique_cnt != len(nums) {
			unique_cnt++
		}
		red = unique_cnt
		if nums[red] == val && red != len(nums) {
			red++
		}
		if unique_cnt < len(nums) && red < len(nums) {
			nums[unique_cnt], nums[red] = nums[red], nums[unique_cnt]
		}
	}
	return unique_cnt
}