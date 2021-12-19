package main

import "fmt"

type stack struct {
	mins []int
	data []int
}

func (s *stack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

func (s *stack) Push(item int) {
	s.data = append(s.data, item)
	if len(s.mins) == 0 || item <= s.GetMin() {
		s.mins = append(s.mins, item)
	}
}

func (s *stack) Pop() int {
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if s.GetMin() == ret {
		s.mins = s.mins[:len(s.mins)-1]
	}
	return ret
}

func main() {
	s := stack{}
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	fmt.Println(s.Pop())
	fmt.Println(s.GetMin())
}
