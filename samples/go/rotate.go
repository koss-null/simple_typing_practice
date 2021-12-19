package main

import "fmt"

var changed map[[2]int]struct{}

func change(matrix [][]int, val int, i, j int) {
	if _, ok := changed[[2]int{i, j}]; ok {
		return
	}
	mem := matrix[len(matrix)-j-1][len(matrix)-i-1]
	matrix[len(matrix)-j-1][len(matrix)-i-1] = val
	changed[[2]int{i, j}] = struct{}{}
	change(matrix, mem, len(matrix)-j-1, len(matrix)-i-1)

}

func rotate(matrix [][]int) {
	changed = map[[2]int]struct{}{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			change(matrix, matrix[i][j], i, j)
		}
	}
}

func main() {
	fmt.Println("vim-go")
}
