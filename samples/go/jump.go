package main

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	lf, rg := 0, 0
	for k := 0; k <= len(nums); k++ {
		maxDist := 0
		for i := lf; i <= rg; i++ {
			if nums[i]+i+1 >= len(nums) {
				return k + 1
			}
			if maxDist < i+nums[i] {
				maxDist = i + nums[i]
			}
		}
		lf = rg + 1
		rg = maxDist
	}
	return 1
}

func main() {
	//fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	// //fmt.Println(jump([]int{2, 0, 8, 0, 3, 4, 7, 5, 6, 1, 0, 0, 5, 9, 7, 5, 3, 6}))
}
