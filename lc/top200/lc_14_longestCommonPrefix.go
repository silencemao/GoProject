package main

import (
	"fmt"
	"math"
)

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		tmp, j := strs[i], 0
		for ; j < int(math.Min(float64(len(res)), float64(len(tmp)))); j++ {
			if tmp[j] != res[j] {
				break
			}
		}
		res = res[:j]
	}
	return res
}
func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
