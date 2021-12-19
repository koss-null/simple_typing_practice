package main

import "fmt"

func splitIdx(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	split_idx := splitIdx(nums)
	fmt.Println(split_idx)
	var lf, rg int
	if split_idx == -1 {
		lf, rg = 0, len(nums)-1
	} else if target > nums[len(nums)-1] {
		lf, rg = 0, split_idx
	} else {
		lf, rg = split_idx+1, len(nums)-1
	}
	fmt.Println("lfrg", lf, rg)

	for lf <= rg {
		md := (rg-lf)/2 + lf
		if nums[md] <= target {
			lf = md + 1
		} else {
			rg = md - 1
		}
	}
	fmt.Println("lfrg", lf, rg)

	if rg > -1 && nums[rg] == target {
		return rg
	}
	if lf < len(nums) && nums[lf] == target {
		return lf
	}
	return -1
}

func main() {
	fmt.Println(search([]int{3, 5, 1}, 3))
}
