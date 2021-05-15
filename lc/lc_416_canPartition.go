package main

import "fmt"

func canPartition(nums []int) bool {
	if len(nums) < 1{
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum / 2
	if (sum & 1)==1 {
		return false
	}
	dp := make([]bool, target + 1)
	dp[0] = true

	for _, num := range nums {
		for j := target; j > 0; j-- {
			if j >= num {
				dp[j] = dp[j] || dp[j - num]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}