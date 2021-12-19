package main

import "fmt"

type pair struct {
	val   float64
	power int
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.
	}

	minus := false
	if n < 0 {
		n *= -1
		minus = true
	}

	powers := make([]pair, 0) // powers of 2 starting from 1
	res, power := x, 1
	for n > power*2 {
		res *= res
		power *= 2
		powers = append(powers, pair{res, power})
	}

	for i := len(powers) - 1; i > -1; i-- {
		for power+powers[i].power <= n {
			res *= powers[i].val
			power += powers[i].power
		}
	}

	if minus {
		return 1. / res
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
