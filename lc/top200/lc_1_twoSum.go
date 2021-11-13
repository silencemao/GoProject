package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	set := make(map[int]int, 0)
	for ind1, num := range nums {
		if ind, ok := set[target-num]; ok {
			res = []int{ind, ind1}
			return res
		} else {
			set[num] = ind1
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
