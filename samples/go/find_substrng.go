package main

import (
	"fmt"
)

func incert(idxs map[string][]int, s string, i int) map[string][]int {
	if a, ok := idxs[s]; ok {
		idxs[s] = append(a, i)
	} else {
		idxs[s] = []int{i}
	}
	return idxs
}

func check_idx(used_idxs []int, idx int, word_length int) bool {
	for _, i := range used_idxs {
		if idx >= i && idx < i+word_length {
			return false
		}
	}
	return true
}

func peak(first_word_idxs map[string][]int, words []string, used_idxs []int, word_length int) []int {
	if len(words) == 0 {
		fmt.Println("the len of words is 0, finishing", used_idxs)
		return used_idxs
	}

	first_idxs, ok := first_word_idxs[words[0]]
	fmt.Printf("the word is %s, idxs are %+v \n", words[0], first_idxs)
	if !ok {
		return []int{}
	}
	for _, wi := range first_idxs {
		fmt.Printf("%d checking it\n", wi)
		if check_idx(used_idxs, wi, word_length) {
			fmt.Println("it's ok")
			new_words := words[0:wi]
			if wi+1 < len(words) {
				new_words = append(new_words, words[wi+1:]...)
			}
			fmt.Println("new words are: ", new_words)

			res := peak(
				first_word_idxs,
				new_words,
				append(used_idxs, wi),
				word_length,
			)
			fmt.Println("the res is: ", res)
			if len(res) == len(words) {
				return res
			}
		}
	}

	return []int{}
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	start_indexes := make(map[string][]int)
	word_len := len(words[0])
	for i := range s {
		if i+word_len < len(s) {
			start_indexes = incert(start_indexes, s[i:i+word_len], i)
		}
	}
	fmt.Println(start_indexes)

	res := peak(start_indexes, words, make([]int, 0), word_len)
	if len(res) == len(words) {
		return res
	}
	return []int{}

}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
