package main

import "fmt"

var charToNum = map[byte]byte{
	'A': 1,
	'B': 2,
	'C': 3,
	'D': 4,
	'E': 5,
	'F': 6,
	'G': 7,
	'H': 8,
	'I': 9,
	'J': 10,
	'K': 11,
	'L': 12,
	'M': 13,
	'N': 14,
	'O': 15,
	'P': 16,
	'Q': 17,
	'R': 18,
	'S': 19,
	'T': 20,
	'U': 21,
	'V': 22,
	'W': 23,
	'X': 24,
	'Y': 25,
	'Z': 26,
}

func titleToNumber(columnTitle string) int {
	res := 0
	mem := 1
	for i := len(columnTitle) - 1; i > -1; i-- {
		res += int(charToNum[columnTitle[i]]) * mem
		mem *= 26
	}
	return res
}

func main() {
	fmt.Println(titleToNumber("Y"))
}
