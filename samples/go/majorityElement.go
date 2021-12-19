package main

import "fmt"

func majorityElement(nums []int) int {
	counts := map[int]int{}
	for i := range nums {
		counts[nums[i]] += 1
		if counts[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}

func main() {
	fmt.Println("vim-go")
}
