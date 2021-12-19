package main

// stolen one

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	was1 := false
	n := len(nums)
	for i := range nums {
		if nums[i] == 1 {
			was1 = true
		}
		if nums[i] >= len(nums) || nums[i] < 0 {
			nums[i] = 0
		}
	}
	if !was1 {
		return 1
	}

	for i := range nums {
		//fmt.Println(nums)
		nums[nums[i]%n] += n
	}
	//fmt.Println(nums)
	for i := range nums {
		//fmt.Println("check", nums[i], len(nums), nums[i] / len(nums))
		if nums[i]/len(nums) == 0 {
			return i
		}
	}
	return n
}
