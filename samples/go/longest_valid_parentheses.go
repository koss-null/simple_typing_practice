package main

func lenOfValidParentheses(start int, s string) int {
	if s[start] == ')' {
		return 0
	}
	op_br, cl_br := 1, 0
	cnt := 1
	for op_br > cl_br && cnt+start < len(s) {
		switch s[cnt+start] {
		case '(':
			op_br++
		case ')':
			cl_br++
		}
		cnt++
	}
	if op_br == cl_br {
		return cnt
	}
	return 0
}

func longestValidParentheses(s string) int {
	par_len := make([]int, len(s))
	for i := 0; i < len(s); {
		par_len[i] = lenOfValidParentheses(i, s)
		if par_len[i] == 0 {
			i += 1
		} else {
			i += par_len[i]
		}
	}

	max, cur, cur_i := 0, 0, 0
	for cur_i < len(par_len) {
		if par_len[cur_i] != 0 {
			cur += par_len[cur_i]
			cur_i += par_len[cur_i]
		} else {
			if cur > max {
				max = cur
			}
			cur = 0
			cur_i++
		}
	}
	if cur > max {
		max = cur
	}
	return max
}
