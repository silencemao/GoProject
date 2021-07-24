package main

import "fmt"

/*
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
方案一：遍历每个数字 遍历数字的每一位 计数
方案二：
因为当前数i比前一个数字i-1大1，所以相当于在i-1的最后一位加1。分两种情况讨论，看是否存在进位，
若i-1的最后一位为0，则i的位数是i-1的位数+1，因为不存在进位。
若i-1的最后一位是1，则存在进位。我们先利用一个变量tmp记录i-1的位数，然后遍历i-1的每一位，若是1，则tmp-1，因为会进位，若是0，则break出来，进位结束tmp+1
*/
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if (i-1)&1 == 1 {
			m := i - 1
			tmp := res[i-1]
			for m > 0 && m&1 == 1 {
				tmp -= 1
				m >>= 1
			}
			tmp += 1
			res[i] = tmp
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}

func countBits1(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&1 == 1 { // 奇数比前一个偶数多1
			res[i] = res[i-1] + 1
		} else { // 偶数相当于i/2向左移动一位，1的个数不变
			res[i] = res[i/2]
		}
	}
	return res
}

func main() {
	n := 7

	fmt.Println(countBits(n))
	fmt.Println(countBits1(7))
}
