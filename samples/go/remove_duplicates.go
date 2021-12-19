
func removeDuplicates(nums []int) int {
	green, red := 0, 0
	unique_cnt := 0
	for {
		red++
		if red == len(nums) {
			break
		}
		if nums[red] != nums[green] {
			unique_cnt++
			nums[unique_cnt] = nums[red]
			green = red
		}
	}
	return unique_cnt
}