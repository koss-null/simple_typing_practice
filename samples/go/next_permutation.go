package main

import "fmt"

func getMinGreater(target int, nums []int) int {
	lf, rg := 0, len(nums)-1
	for rg >= lf {
		md := (rg-lf)/2 + lf
		if nums[md] > target {
			lf = md + 1
		} else {
			rg = md - 1
		}
	}
	return rg
}

func reverse(nums []int, start_idx int) {
	lf, rg := start_idx, len(nums)-1
	for lf < rg {
		nums[lf], nums[rg] = nums[rg], nums[lf]
		lf++
		rg--
	}
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i != -1 {
		min_greater_idx := getMinGreater(nums[i], nums[i+1:]) + i + 1
		nums[i], nums[min_greater_idx] = nums[min_greater_idx], nums[i]
	}
	reverse(nums, i+1)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
	nextPermutation(a)
	fmt.Println(a)
}
