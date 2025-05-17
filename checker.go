package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	word := "Madam"
	if isPalindrome(word) {
		fmt.Println(word, "is a palindrome!")
	} else {
		fmt.Println(word, "is not a palindrome.")
	}
}
