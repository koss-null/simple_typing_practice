package main

import "bytes"

func isPalindrome(s string) bool {
	trimmed := make([]byte, 0, len(s))
	for i := range s {
		if (s[i] >= 'a' && s[i] <= 'z') ||
			(s[i] >= 'A' && s[i] <= 'Z') ||
			(s[i] >= '0' && s[i] <= '9') {
			trimmed = append(trimmed, s[i])
		}
	}
	trimmed = bytes.ToLower(trimmed)
	for i := 0; i < len(trimmed)/2; i++ {
		if trimmed[i] != trimmed[len(trimmed)-i-1] {
			return false
		}
	}
	return true
}
