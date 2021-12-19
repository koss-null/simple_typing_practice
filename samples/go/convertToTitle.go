package main

import "fmt"

var charToNum = []byte{
	'Z',
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'H',
	'I',
	'J',
	'K',
	'L',
	'M',
	'N',
	'O',
	'P',
	'Q',
	'R',
	'S',
	'T',
	'U',
	'V',
	'W',
	'X',
	'Y',
}

const alphabetLen = 26

func convertToTitle(columnNumber int) string {
	res := make([]byte, 0, 64)
	for columnNumber != 0 {
		// fmt.Println("changing res", byte((columnNumber % alphabetLen)))
		res = append(res, charToNum[columnNumber%alphabetLen])
		cns := columnNumber
		columnNumber /= alphabetLen
		if cns%alphabetLen == 0 {
			columnNumber--
		}
		// fmt.Println(res)
	}

	for i := 0; i < len(res)/2; i++ {
		// fmt.Println("changing")
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(convertToTitle(2147483647))
}
