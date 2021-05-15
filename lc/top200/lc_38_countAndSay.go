package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	content := "1"
	for i := 1; i < n; i++ {
		content = count(content)
	}
	return content
}

func count(content string) string {
	content += " "
	l := 0
	times := 0
	res := ""
	for r := 1; r < len(content); {
		if content[r] != content[l] {
			times = r - l
			res += strconv.Itoa(times) + string(content[l])
			l = r
		}
		r++
	}
	return res
}

func main() {
	fmt.Println(countAndSay(4))
}
