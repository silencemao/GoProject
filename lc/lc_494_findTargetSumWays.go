package main

import "fmt"

/*
给定一个数组nums， 一个target S
可以对数组nums的数字任意使用+ -号，使得其结果等于S
有多少种方法
*/
func sumWaysHelp(nums []int, S int, index int, sum int, res *int) {
	if index == len(nums){
		if sum == S {
			*res++
		}
		return
	}
	sumWaysHelp(nums, S, index+1, sum + nums[index], res)
	sumWaysHelp(nums, S, index+1, sum - nums[index], res)

	//for i := index; i < len(nums); i++ {
	//	*sum = *sum + nums[i]
	//	sumWaysHelp(nums, S, i+1, sum, res)
	//	*sum = *sum - nums[i]
	//}
}

func findTargetSumWays(nums []int, S int) int {
	sum, res := 0, 0
	sumWaysHelp(nums, S, 0, sum, &res)
	return res
}

/*
                  sum(p) - sum(n) = target
sum(p) + sum(n) + sum(p) - sum(n) = sum(p) + sum(n) + target
                          2sum(p) = sum(nums) + target
sum(p) = (sum(nums) + target) / 2
*/

func findTargetSumWays1(nums []int, S int) int {
	res := 0

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if (sum + S) % 2 == 0 {
		target := (sum + S) / 2
	    dp := make([]int, target+1)
	    dp[0] = 1
	    for _, n := range nums {
	    	for i := target; i >= n; i-- {
	    		dp[i] += dp[i - n]
			}
		}
		res = dp[target]
	}

	return res
}


func main() {

	nums := []int{1,2,7,9,981}
	S := 1000000000
	fmt.Println(findTargetSumWays(nums, S))
	fmt.Println(findTargetSumWays1(nums, S))
}
