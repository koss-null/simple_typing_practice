import "math"

func pow(k, n int) int {
	if n == 0 {
		return 1
	}
	res := 0
	for i := 0; i < n; i++ {
		res += k
	}
	return res
}

func divide(dividend int, divisor int) int {
	negative := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		negative = true
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	divisor_mult := make([]int, 0, 31)
	divisor_mult = append(divisor_mult, divisor)
	divisor_amt := make([]int, 0, 31)
	for i := 0; i < 32; i++ {
		divisor_amt = append(divisor_amt, pow(2, i))
	}

	for divisor_mult[len(divisor_mult)-1] < dividend {
		last_dm := divisor_mult[len(divisor_mult)-1]
		divisor_mult = append(divisor_mult, last_dm+last_dm)
	}

	res := 0
	for i := len(divisor_mult) - 2; i >= 0; i-- {
		for dividend-divisor_mult[i] >= 0 {
			dividend -= divisor_mult[i]
			res += divisor_amt[i]
		}
	}
	if negative {
		return -res
	}
	return res
}