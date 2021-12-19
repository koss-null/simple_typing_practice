package main

func binSearch(data [][]int, target int) bool {
	if len(data) == 0 || len(data[0]) == 0 {
		return false
	}

	M, N := len(data), len(data[0])
	lf, rg := 0, (M*N)-1
	for lf <= rg {
		md := (rg-lf)/2 + lf
		m := md / N
		n := N - (md - m*N) - 1
		if data[m][n] <= target {
			rg = md - 1
		} else {
			lf = md + 1
		}
	}

	rg += 1
	m := rg / N
	n := N - (rg - m*N) - 1
	return data[m][n] == target
}

func main() {
	// fmt.Println(binSearch([][]int{{1}}, 2))
	// fmt.Println(binSearch([][]int{{2}}, 2))
	// fmt.Println(binSearch([][]int{{1, 3, 4, 5}}, 2))
	// fmt.Println(binSearch([][]int{
	// 	{1, 3, 4, 5},
	// 	{6, 7, 9, 10},
	// }, 2))
	// fmt.Println(binSearch([][]int{
	// 	{1, 2, 4, 5},
	// 	{6, 7, 9, 10},
	// }, 2))
	// fmt.Println(binSearch([][]int{{}, {}}, 0))
}
