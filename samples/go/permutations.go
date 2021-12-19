package main

import (
	"container/list"
	"fmt"
	//"fmt"
)

func getSlice(l *list.List) []int {
	res := make([]int, 0, l.Len())
	for el := l.Front(); el != nil; el = el.Next() {
		res = append(res, el.Value.(int))
	}
	return res
}

func getKeys(m map[int]int) []int {
	res := make([]int, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func getPerms(val int, available map[int]int, before *list.List) [][]int {
	//fmt.Println("checking val ", val, available, getSlice(before))
	if len(available) == 0 {
		//fmt.Println("finish branch")
		return [][]int{getSlice(before)}
	}

	res := make([][]int, 0)
	keys := getKeys(available)
	for _, k := range keys {
		//fmt.Println("going to branch", k)
		available[k]--
		if available[k] == 0 {
			delete(available, k)
		}
		before.PushBack(k)

		res = append(res, getPerms(k, available, before)...)

		before.Remove(before.Back())
		available[k]++
	}
	return res
}

func permute(nums []int) [][]int {
	available := make(map[int]int, len(nums))
	for i := range nums {
		available[nums[i]]++
	}

	res := make([][]int, 0)
	for i := range nums {
		available[nums[i]]--
		if available[nums[i]] == 0 {
			delete(available, nums[i])
		}
		before := list.New()
		before.PushBack(nums[i])
		res = append(res, getPerms(nums[i], available, before)...)
		available[nums[i]]++
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
