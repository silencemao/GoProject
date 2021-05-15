package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func jump(nums []int) int {
	res, len := 0, len(nums) - 1
	pre, cur, i := 0, 0, 0

	for cur < len {
		res ++
		pre = cur
		for ; i <= pre; i++ {
			cur = max(cur, i + nums[i])
		}
		if pre==cur {
			return -1
		}
	}
	return res
}

/*
给定一个数组nums，每一个数字都代表在当前位置能跳的最远距离，返回当前数组跳到最后位置所需要的步数
贪心算法，本题的贪心并不是每一步都跳最大的步数，而是在当前能跳力范围内下一步能跳的最远范围即i+num[i]

last代表上一步能到达的最远范围，cur代表当前能到达的最远范围，当i到达上一步能到达的最远范围时，说明还没有到达最后一个位置，此时更新
last以及cur
*/
func jump1(nums []int) int {
	res, len := 0, len(nums) - 1
	last, cur := 0, 0
	for i := 0; i < len; i++ {
		cur = max(cur, i + nums[i])
		if i==last {
			res++
			last = cur
			if cur >= len {
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums), jump1(nums))
}
