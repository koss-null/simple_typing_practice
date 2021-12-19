package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	lf, rg := 0, len(nums)-1
	for lf <= rg {
		md := (rg-lf)/2 + lf
		if nums[md] < target {
			lf = md + 1
		} else {
			rg = md - 1
		}
	}
	targetStart := rg + 1
	if targetStart >= len(nums) || nums[targetStart] != target {
		return []int{-1, -1}
	}

	lf, rg = 0, len(nums)-1
	for lf <= rg {
		md := (rg-lf)/2 + lf
		if nums[md] <= target {
			lf = md + 1
		} else {
			rg = md - 1
		}
	}
	return []int{targetStart, lf - 1}
}
