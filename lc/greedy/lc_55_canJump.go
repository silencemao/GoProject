package main

import "GoProject/leetcode/util"

func canJump(nums []int) bool {
	k := 0 // 能到达的最远距离
	for i := 0; i < len(nums); i++ {
		if k < i { // 最远都到不了i，肯定无法到达尾部
			return false
		}
		k = util.MaxInt(nums[i]+i, k) // 当前位置能到达的最远距离
		if k >= len(nums)-1 {
			return true
		}
	}
	return true
}
func main() {

}
