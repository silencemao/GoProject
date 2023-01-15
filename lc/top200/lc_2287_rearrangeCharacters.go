package main

import "fmt"

func rearrangeCharacters(s string, target string) int {
	res := 1<<31 - 1

	set1 := map[int]int{}
	set2 := map[int]int{}
	for _, c := range target {
		set1[int(c-'a')] += 1
	}
	for _, c := range s {
		set2[int(c-'a')] += 1
	}
	for k := range set1 {
		if n, ok := set2[k]; ok {
			if n/set1[k] < res {
				res = n / set1[k]
			}
		} else {
			return 0
		}
	}
	return res
}

func main() {
	s := "abbaccaddaeea"
	target := "aaaaa"
	fmt.Println(rearrangeCharacters(s, target))
}
