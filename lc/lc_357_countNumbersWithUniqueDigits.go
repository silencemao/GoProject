package main

import "fmt"

/*
给定数字n，返回0-10^n之间 有多少个 数是无重复数字组成的
如n=1 返回10
  n=2 返回91

最高位有9种选择，此高位也有9种选择，向后依次类推
我们可以分为0-9、10-99、100-999...这样按照位数来划分
计算不同范围内的数字个数的时候是互不影响的，记住 互不影响的，我就是由于这一点没有想清楚，导致结果一直不对

*/
/*
直接两重循环，最外层循环表示当前的数字一共多少位，最高位的数字有9种选择，最内层循环表示除最高位之外，其余的每一位数字有多少种选择
*/
func countNumbersWithUniqueDigits(n int) int {
	if n==0 {
		return 1
	}
	res := 10

	for i := 2; i <= n; i++ {
		base := 9
		for j := 0; j < i - 1; j++ {
			base *= 9 - j
		}
		res += base
	}

	return res
}

func helpCNWD(n, i int) int {
	if i==0 {
		return 1
	} else {
		return n * helpCNWD(n - 1, i - 1)
	}
}

/*
采用递归的方式实现，思路和上面差不多，也是按照位数来划分范围，然后计算每一个范围内的个数
*/
func countNumbersWithUniqueDigits1(n int) int {
	res := 1

	for i := 0; i < n; i++ {
		res += 9 * helpCNWD(9, i)
	}

	return res
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(4))
	fmt.Println(countNumbersWithUniqueDigits1(4))
}
