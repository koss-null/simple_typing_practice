package main

import "fmt"

type (
	RegItem interface {
		Matches(s string, i int) func() (bool, int)
	}

	char struct {
		val byte
	}

	any struct { /* ? */
	}

	anyN struct { /* * */
	}
)

func (c *char) Matches(s string, i int) func() (bool, int) {
	if i < len(s) && s[i] == c.val {
		first := true
		return func() (bool, int) {
			if first {
				first = false
				return true, i + 1
			}
			return false, -1
		}
	}
	return nil
}

func (_ *any) Matches(s string, i int) func() (bool, int) {
	if i < len(s) {
		first := true
		return func() (bool, int) {
			if first {
				first = false
				return true, i + 1
			}
			return false, -1
		}
	}
	return nil
}

func (_ *anyN) Matches(s string, i int) func() (bool, int) {
	var idx = i
	return func() (bool, int) {
		if idx <= len(s) {
			idx++
			return true, idx - 1
		}
		return false, -1
	}
}

var regToPref map[*RegItem]map[int]bool

func putToMap(ri *RegItem, idx int, ans bool) {
	if m, ok := regToPref[ri]; ok {
		m[idx] = ans
		return
	}
	regToPref[ri] = make(map[int]bool)
	regToPref[ri][idx] = ans
}

func matchReg(s string, reg []RegItem, idx int) bool {
	if idx == len(s) {
		if len(reg) == 1 {
			next := reg[0].Matches(s, idx)
			if next == nil {
				return false
			}
			ok, _ := next()
			putToMap(&reg[0], idx, ok)
			return ok
		}
		return len(reg) == 0
	}
	if len(reg) == 0 {
		return false
	}

	if prefToRes, ok := regToPref[&reg[0]]; ok {
		if res, ok := prefToRes[idx]; ok {
			return res
		}
	}

	next := reg[0].Matches(s, idx)
	if next == nil {
		putToMap(&reg[0], idx, false)
		return false
	}

	ok, indx := next()
	for ok {
		if matchReg(s, reg[1:], indx) {
			putToMap(&reg[0], idx, true)
			return true
		}
		ok, indx = next()
	}

	putToMap(&reg[0], idx, false)
	return false
}

func trimAstetisks(p string) string {
	cur := 0
	prev := byte('0')
	res := make([]byte, 0, len(p))
	for cur < len(p) {
		if prev == '*' && p[cur] == '*' {
			cur++
			continue
		}
		prev = p[cur]
		res = append(res, p[cur])
		cur++
	}
	return string(res)
}

func isMatch(s string, p string) bool {
	regToPref = make(map[*RegItem]map[int]bool)

	p = trimAstetisks(p)

	parsedReg := make([]RegItem, 0, len(p))
	for i := range p {
		switch p[i] {
		case '?':
			parsedReg = append(parsedReg, &any{})
		case '*':
			parsedReg = append(parsedReg, &anyN{})
		default:
			parsedReg = append(parsedReg, &char{p[i]})
		}
	}

	return matchReg(s, parsedReg, 0)
}

func main() {
	fmt.Println(isMatch("aa", "*"))
}
