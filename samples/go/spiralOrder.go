package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(matrix)*len(matrix[0]))
	usedTop, usedBottom, usedLeft, usedRight := 0, 0, 0, 0
	totalLength := len(matrix) * len(matrix[0])
	for len(result) < totalLength {
		for i := usedLeft; i < len(matrix[0])-usedRight; i++ {
			result = append(result, matrix[usedTop][i])
			//fmt.Println(matrix[usedTop][i], ", ")
		}
		usedTop++
		if len(result) >= totalLength {
			break
		}
		for i := usedTop; i < len(matrix)-usedBottom; i++ {
			result = append(result, matrix[i][len(matrix[usedTop])-usedRight-1])
			//fmt.Println(matrix[i][len(matrix[usedTop])-usedRight-1], ", ")
		}
		usedRight++
		if len(result) >= totalLength {
			break
		}
		for i := len(matrix[0]) - usedRight - 1; i >= usedLeft; i-- {
			result = append(result, matrix[len(matrix)-usedBottom-1][i])
			//fmt.Println(matrix[len(matrix)-usedBottom-1][i], ", ")
		}
		usedBottom++
		if len(result) >= totalLength {
			break
		}
		for i := len(matrix) - usedBottom - 1; i >= usedTop; i-- {
			result = append(result, matrix[i][usedLeft])
			//fmt.Println(matrix[i][usedLeft], ", ")
		}
		usedLeft++
	}
	return result
}
