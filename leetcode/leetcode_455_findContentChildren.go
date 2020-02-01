package main

import (
	"fmt"
	"sort"
)
/*
给定一个数组g表示小朋友的满意度
给定一个数组s表示每一个物品的价，将物品分给g中的小朋友，若价值大于小朋友的满意度，则表示分配成功，
计算最大的满意度
贪心算法，遍历s如果s[i]，在g中找到第一个满足条件的g[j]，然后res++，并去除g[j]
*/
func findContentChildren(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(g); j++ {
			if s[i] >= g[j] {
				res++
				g = append(g[:j], g[j+1:]...)
				break
			}
		}
	}
	return res
}

func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for i := 0; i < len(s) && child < len(g); i++ {
		if s[i] >= g[child] {
			child++
		}
	}
	return child
}

func main() {
	g := []int{1, 2, 3, 4, 5}
	s := []int{1, 1, 4, 5, 5}
	//fmt.Println(findContentChildren(g, s))
	fmt.Println(findContentChildren1(g, s))
}
