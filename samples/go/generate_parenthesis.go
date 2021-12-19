package main

import (
	"fmt"
)

func getNextValidParenthesis(sequence []byte, op_b, cl_b, n int, next_p byte) []string {
	sequence = append(sequence, next_p)
	if op_b == n && cl_b == op_b {
		return []string{string(sequence)}
	}
	var res1, res2 []string
	if cl_b < op_b {
		res1 = getNextValidParenthesis(sequence, op_b, cl_b+1, n, ')')
	}
	if op_b < n {
		res2 = getNextValidParenthesis(sequence, op_b+1, cl_b, n, '(')
	}
	return append(res1, res2...)
}

func generateParenthesis(n int) []string {
	return getNextValidParenthesis(make([]byte, 0, 10), 1, 0, n, '(')
}

func main() {
	fmt.Println(generateParenthesis(3))
}
