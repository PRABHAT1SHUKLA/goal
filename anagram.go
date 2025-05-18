package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(a, b string) bool {
	a = strings.ToLower(strings.ReplaceAll(a, " ", ""))
	b = strings.ToLower(strings.ReplaceAll(b, " ", ""))
	
	aSlice := strings.Split(a, "")
	bSlice := strings.Split(b, "")
	
	sort.Strings(aSlice)
	sort.Strings(bSlice)
	
	return strings.Join(aSlice, "") == strings.Join(bSlice, "")
}

func main() {
	fmt.Println(isAnagram("Listen", "Silent")) // true
	fmt.Println(isAnagram("Hello", "Olelh"))   // true
}
