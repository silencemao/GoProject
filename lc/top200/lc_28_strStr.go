package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	n := len(needle)
	for i := 0; i <= len(haystack)-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack, needle := "hello", "llo"
	fmt.Println(strStr(haystack, needle))
}
