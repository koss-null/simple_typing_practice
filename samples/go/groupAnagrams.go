package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))
	for i := range strs {
		runeStr := make([]rune, 0, len(strs[i]))
		for _, c := range strs[i] {
			runeStr = append(runeStr, c)
		}

		sort.Slice(runeStr, func(i, j int) bool { return runeStr[i] < runeStr[j] })
		anagramKey := string(runeStr)

		if _, ok := anagrams[anagramKey]; ok {
			anagrams[anagramKey] = append(anagrams[anagramKey], strs[i])
			continue
		}
		anagrams[anagramKey] = []string{strs[i]}
	}

	res := make([][]string, 0, len(anagrams))
	for k := range anagrams {
		res = append(res, anagrams[k])
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
