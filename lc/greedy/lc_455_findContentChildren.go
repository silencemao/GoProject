package main

import (
	"fmt"
	"sort"
)

/*
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j]。如果 s[j]>= g[i]，
我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

贪心思路
*/
func findContentChildren(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(g); i++ {
		j := 0
		for ; j < len(s); j++ {
			if s[j] >= g[i] {
				res += 1
				s = append(s[:j], s[j+1:]...)
				break
			}
		}

	}
	return res
}

func findContentChildren1(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)
	ind := len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		if ind >= 0 && s[ind] >= g[i] {
			res += 1
			ind--
		}
	}

	return res
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
	fmt.Println(findContentChildren1(g, s))
}
