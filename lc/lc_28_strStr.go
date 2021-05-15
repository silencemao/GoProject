package main

import "fmt"

func strStr(haystack string, needle string) int {
	res := -1
	if haystack==needle || len(needle)==0 {
		return 0
	}
	if len(needle) > len(haystack) || len(haystack)==0  {
		return res
	}

	for i := 0; i <= len(haystack) - len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return res
}

func main() {
	haystack := "aaaaa"
	needle := "bba"
	fmt.Println(strStr(haystack, needle))
}

