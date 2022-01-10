package main

import "fmt"

/*
1、如果是问填满背包，则递推式为 dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
如果是问填满背包有多少个方法，则为dp[j] += dp[j-coins[i]]
2、如果是01背包问题，则物品只能用一次，j必须从后向前遍历 for j:=amount;j>=coins[i];j--
如果是完全背包问题，则物品能用无数次，j必须从前向后遍历 for j:= coins[i] ; j <= amount ;j++
3、 先循环物品还是先循环背包容量？
如果懒得判断的话就直接先物品再背包容量，99%对。
A:01背包中二维dp数组的两个for遍历的先后循序是可以颠倒了，一维dp数组的两个for循环先后循序一定是先遍历物品，再遍历背包容量（否则前面的值会被覆盖）。
B:在纯完全背包中，对于一维dp数组来说，其实两个for循环嵌套顺序同样无所谓！因为dp[j] 是根据 下标j之前所对应的dp[j]计算出来的。 只要保证下标j之前的dp[j]都是经过计算的就可以了。有些问多少中方法填满的变种题，i,j的顺序涉及到组合排列--物品,容量为组合。容量,物品为排列。
4、dp[0]初始化。问填满背包的方法，dp[0]=1;问能否填满背包，dp[0]=0


给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 凑成总金额为0的货币组合数为1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c] // 考虑货币c的组合数 dp[i-c] + 不考虑货币c的组合数dp[i]
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(change(amount, coins))
}
