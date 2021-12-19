package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	lf, rg := 0, len(numbers)-1
	for numbers[lf]+numbers[rg] != target {
		if numbers[lf]+numbers[rg] > target {
			rg--
		} else {
			lf++
		}
	}
	return []int{lf + 1, rg + 1}
}

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
