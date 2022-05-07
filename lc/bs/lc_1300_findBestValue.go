package main

import (
	"fmt"
	"sort"
)

func findBestValue(arr []int, target int) int {
	n, pre := len(arr), 0
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		tmp := pre + (n-i)*arr[i]
		if tmp >= target {
			value := (target - pre) / (n - i)
			sumLow := pre + value*(n-i)
			sumHigh := pre + (value+1)*(n-i)
			if target-sumLow <= sumHigh-target {
				return value
			} else {
				return value + 1
			}
		}
		pre += arr[i]
	}
	return arr[len(arr)-1]
}

func main() {
	arr := []int{2, 3, 5}
	target := 10
	fmt.Println(findBestValue(arr, target))
}
