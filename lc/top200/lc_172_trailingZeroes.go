package main

import "fmt"

/*
给定数字n，计算n！之后数字尾部0的个数
尾部0的个数，其实就是计算有多少10，每个10的最小因子是2和5，而2及2的倍数的数字明显要多于5及5的倍数
因此最后0的个数是由5的个数决定，所以我们计算0-n以内一共有多少个5
此处多少个5并不仅仅是n/5，而是1-n之间每个数字进行拆分之后有多少个5
n=10
5 10=2*5
1  1
*/
func trailingZeroes(n int) int {
	res := 0
	for i := 5; i <= n; i += 5 {
		tmp := i
		for tmp%5 == 0 && tmp >= 5 {
			res += 1
			tmp = tmp / 5
		}
	}
	return res
}

func trailingZeroes1(n int) int {
	res := 0
	for n > 5 {
		res += n / 5
		n = n / 5
	}
	return res
}

func main() {
	n := 10
	fmt.Println(trailingZeroes(n))
	fmt.Println(trailingZeroes1(n))
}
