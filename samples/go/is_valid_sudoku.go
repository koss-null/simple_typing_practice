package main

func validateRows(board [][]byte) bool {
	for row := range board {
		used := make(map[byte]interface{}, 10)
		for i := range board[row] {
			if board[row][i] != '.' {
				if _, ok := used[board[row][i]-byte('0')]; ok {
					return false
				}
				used[board[row][i]-byte('0')] = struct{}{}
			}
		}
	}

	return true
}

func validateColumns(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		used := make(map[byte]interface{}, 10)
		for row := 0; row < len(board); row++ {
			if board[row][i] != '.' {
				if _, ok := used[board[row][i]-byte('0')]; ok {
					return false
				}
			}
			used[board[row][i]-byte('0')] = struct{}{}
		}
	}

	return true
}

func validateBoxes(board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			used := make(map[byte]interface{}, 10)
			for line := 3 * i; line < 3*(i+1); line++ {
				for row := 3 * j; row < 3*(j+1); row++ {
					if board[line][row] != '.' {
						if _, ok := used[board[line][row]-byte('0')]; ok {
							return false
						}
					}
					used[board[line][row]-byte('0')] = struct{}{}
				}
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	return validateRows(board) && validateColumns(board) && validateBoxes(board)
}
