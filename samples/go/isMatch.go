package main

import (
	"container/list"
)

type (
	Symbol interface {
		FindMatches(s string) []string
	}

	symbol struct {
		val byte
	}

	anySymbol struct{}

	anyNSymbols struct{}
)

func (sym *symbol) FindMatches(s string) []string {
	if len(s) == 0 {
		return nil
	}
	if sym.val == s[0] {
		//fmt.Println("found match with val ", sym.val)
		return []string{s[1:]}
	}
	return nil
}

func (sym *anySymbol) FindMatches(s string) []string {
	if len(s) == 0 {
		return nil
	}
	//fmt.Println("found match with any simble")
	return []string{s[1:]}
}

func (sym *anyNSymbols) FindMatches(s string) []string {
	res := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, s[i:])
	}
	res = append(res, "")
	return res
}

func checkVariants(variants []string, regList *list.List) bool {
	if regList.Front() == nil {
		for i := range variants {
			//fmt.Println(len(variants[i]))
			if len(variants[i]) == 0 {
				return true
			}
		}
		return false
	}
	el := regList.Front().Value

	//fmt.Printf("el type %T \n", el)
	var elem Symbol
	switch el.(type) {
	case *anySymbol:
		elem = el.(*anySymbol)
	case *symbol:
		elem = el.(*symbol)
	case *anyNSymbols:
		elem = el.(*anyNSymbols)
	}

	//fmt.Println("checking variants:")
	//fmt.Println()

	//fmt.Printf("%T \n", elem)
	for i := range variants {
		//fmt.Println("checking ", variants[i])
		//fmt.Println(elem == nil)
		newVariants := elem.FindMatches(variants[i])
		if newVariants != nil {
			mem := regList.Front()
			regList.Remove(mem)
			//fmt.Println("checking variants for this match")
			if checkVariants(newVariants, regList) {
				return true
			}
			regList.PushFront(mem.Value)
		}
	}
	return false
}

func parseRegexp(regList *list.List, s string) {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '?':
			regList.PushBack(&anySymbol{})
		case '*':
			regList.PushBack(&anyNSymbols{})
		default:
			regList.PushBack(&symbol{s[i]})
		}
	}
}

func trimRegexp(p string) string {
	res := make([]byte, 0, len(p))
	last := byte(1)
	for i := range p {
		if last == '*' && p[i] == '*' {
			continue
		}
		res = append(res, p[i])
		last = p[i]
	}
	return string(res)
}

func isMatch(s string, p string) bool {
	p = trimRegexp(p)

	parsedRegexp := list.New()
	parseRegexp(parsedRegexp, p)

	variants := []string{s}
	return checkVariants(variants, parsedRegexp)
}
