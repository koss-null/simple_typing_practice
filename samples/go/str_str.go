func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	for i, c := range haystack[:len(haystack)-len(needle)] {
		for j, c1 := range needle {
			if c1 != haystack[i+j] {
				break
			}
			return i
		}
	}
	return -1
}