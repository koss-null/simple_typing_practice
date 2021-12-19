package main

import "fmt"

func buildHorMap(board [][]byte) []map[byte]struct{} {
	res := make([]map[byte]struct{}, 0, len(board))
	for i := range board {
		used := make(map[byte]struct{}, 9)
		for j := range board[i] {
			if board[i][j] != '.' {
				used[board[i][j]] = struct{}{}
			}
		}
		res = append(res, used)
	}
	return res
}

func buildVertMap(board [][]byte) []map[byte]struct{} {
	res := make([]map[byte]struct{}, 0, len(board))
	for i := range board {
		used := make(map[byte]struct{})
		for j := range board[i] {
			if board[j][i] != '.' {
				used[board[j][i]] = struct{}{}
			}
		}
		res = append(res, used)
	}
	return res
}

func buildSquareMap(board [][]byte) []map[byte]struct{} {
	res := make([]map[byte]struct{}, 0, len(board))
	for multI := 1; multI < 4; multI++ {
		for i := 3 * (multI - 1); i < 3*multI; i++ {
			used := make(map[byte]struct{}, 9)
			for multJ := 1; multJ < 4; multJ++ {
				for j := 3 * (multJ - 1); j < 3*multJ; j++ {
					if board[j][i] != '.' {
						used[board[j][i]] = struct{}{}
					}
				}
			}
			res = append(res, used)
		}
	}
	return res
}

var avalNums = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func getAval(hMap, vMap, sMap map[byte]struct{}) []byte {
	res := make([]byte, 0, 9)
	for i := range avalNums {
		_, okH := hMap[avalNums[i]]
		_, okV := vMap[avalNums[i]]
		_, okS := sMap[avalNums[i]]
		if !okH && !okV && !okS {
			res = append(res, avalNums[i])
		}
	}
	return res
}

// returns nil if no solution found
func solve(board [][]byte, hMap, vMap, sMap []map[byte]struct{}, iStart, jStart int) [][]byte {
	var i, j int
	for i = iStart; i < len(board); i++ {
		for j = jStart; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				break
			}
		}
		if j < len(board[i]) && board[i][j] == '.' {
			break
		}
		jStart = 0
	}
	if i >= len(board) && j >= len(board[0]) {
		return board
	}

	squareIdx := j/3 + i/3*3
	availableVals := getAval(hMap[i], vMap[j], sMap[squareIdx])
	for _, avVal := range availableVals {
		board[i][j] = avVal

		hMap[i][avVal] = struct{}{}
		vMap[j][avVal] = struct{}{}
		sMap[squareIdx][avVal] = struct{}{}

		solution := solve(board, hMap, vMap, sMap, i, j)
		if solution != nil {
			return solution
		}

		delete(hMap[i], avVal)
		delete(vMap[j], avVal)
		delete(sMap[squareIdx], avVal)
		board[i][j] = '.'
	}
	return nil
}

func solveSudoku(board [][]byte) {
	hMap, vMap, sMap := buildHorMap(board), buildVertMap(board), buildSquareMap(board)
	solve(board, hMap, vMap, sMap, 0, 0)
}

func main() {
	req := [][]byte{
		{'1', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '4', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '5', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '6', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '8', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '9'},
	}
	solveSudoku(req)
	for i := range req {
		fmt.Println(string(req[i]))
	}
}
