package main

import "fmt"

const MODULO = 1000_000_00

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func convert(num string) []int {
	res := make([]int, 0)
	sum := 0
	for i := len(num) - 1; i > -1; i -= 8 {
		for j := max(i-7, 0); j <= i; j++ {
			sum += int(num[j] - '0')
			sum *= 10
		}
		res = append(res, sum/10)
		sum = 0
	}
	if sum != 0 {
		res = append(res, sum/10)
	}
	return res
}

func convertBack(num []int) string {
	res := ""
	for i := len(num) - 1; i > -1; i-- {
		res += fmt.Sprintf("%08d", num[i])
	}
	i := 0
	for i < len(res) && res[i] == '0' {
		i++
	}
	return res[i:]
}

func sum(num1, num2 []int, mem int) []int {
	res := make([]int, 0)
	i := 0
	for ; i < min(len(num1), len(num2)); i++ {
		res = append(res, (mem+num1[i]+num2[i])%MODULO)
		mem = (mem + num1[i] + num2[i]) / MODULO
	}
	for ; i < len(num1); i++ {
		res = append(res, (mem+num1[i])%MODULO)
		mem = (mem + num1[i]) / MODULO
	}
	for ; i < len(num2); i++ {
		res = append(res, (mem+num2[i])%MODULO)
		mem = (mem + num2[i]) / MODULO
	}
	if mem != 0 {
		res = append(res, mem)
	}
	return res
}

func multSingle(num1 []int, num2 int, shift int) []int {
	res := make([]int, 0, len(num1)+shift+1)
	for i := 0; i < shift; i++ {
		res = append(res, 0)
	}
	var mem int
	for i := range num1 {
		res = append(res, (num2*num1[i]+mem)%MODULO)
		mem = (num2*num1[i] + mem) / MODULO
	}
	if mem != 0 {
		res = append(res, mem)
	}
	return res
}

func mult(num1, num2 []int) []int {
	res := make([]int, 0)
	for cur := 0; cur < len(num2); cur++ {
		res = sum(res, multSingle(num1, num2[cur], cur), 0)
	}
	return res
}

func multiply(num1, num2 string) string {
	if (len(num1) == 1 && num1[0] == '0') || (len(num2) == 1 && num2[0] == '0') {
		return "0"
	}
	num1a, num2a := convert(num1), convert(num2)
	fmt.Println(num1a, num2a)
	return convertBack(mult(num1a, num2a))
}

func main() {
	fmt.Println(multiply("123456789012345678", "123456789012345678"))
}
