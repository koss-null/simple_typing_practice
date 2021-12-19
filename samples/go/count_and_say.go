package main

func sayString(s string) string {
	if s == "" {
		return ""
	}

	ss := []byte(s)
	res := make([]byte, 0, len(ss))
	cnt := byte(0)
	prev := ss[0]
	for i := 0; i < len(ss); i++ {
		c := ss[i]
		if prev == c {
			cnt++
		} else {
			res = append(res, '0'+byte(cnt), prev)
			cnt = 0
			i -= 1
			prev = c
		}
	}
	if cnt != 0 {
		res = append(res, '0'+byte(cnt), prev)
	}
	return string(res)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	return sayString(s)
}

func main() {
}
