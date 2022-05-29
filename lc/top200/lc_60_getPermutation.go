package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
第k个排列
*/
func help60(nums []int, ind int, res *[]string) {

	if ind >= len(nums) {
		tmp := ""
		for _, n := range nums {
			tmp += strconv.Itoa(n)
		}
		*res = append(*res, tmp)
	}

	for i := ind; i < len(nums); i++ {
		nums[i], nums[ind] = nums[ind], nums[i]
		help60(nums, ind+1, res)
		nums[i], nums[ind] = nums[ind], nums[i]
	}
}

func getPermutation1(n int, k int) string {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	res := make([]string, 0)
	help60(nums, 0, &res)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	fmt.Println(res[k-1])
	return res[k-1]
}

func help601(nums []int, k *int, res *string, used *[]int, n int) {
	if len(nums) >= n {
		*k -= 1
		if *k == 0 {
			for _, num := range nums {
				*res += strconv.Itoa(num + 1)
			}
			return
		}
	}
	for i := 0; i < n; i++ {
		if (*used)[i] == 0 {
			(*used)[i] = 1
			nums = append(nums, i)
			help601(nums, k, res, used, n)
			(*used)[i] = 0
			nums = nums[:len(nums)-1]

		}
	}
}

func getPermutation(n int, k int) string {
	nums := []int{}
	res := ""
	used := make([]int, n)
	help601(nums, &k, &res, &used, n)

	return res
}

func main() {
	n, k := 3, 4
	fmt.Println(getPermutation(n, k))
	fmt.Println(getPermutation1(n, k))
}
