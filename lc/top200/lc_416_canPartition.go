package main

/*
分割等和子集

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

若数组的和为奇数，直接返回false
若数组的和为偶数，则判断是否存在可以组成和为sum/2的子集
转换成0-1背包问题，逐个遍历数组，当背包中剩余空间i(i>=num)时，判断 若将背包空出num空间时，背包是否装满，即判断dp[i-num] == true

*/
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true

	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}
func main() {

}
