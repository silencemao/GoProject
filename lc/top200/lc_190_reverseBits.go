package main

import "fmt"

/*
反转二进制位
给定一个数字，将其二进制位反转，返回

1、res记录结果
2、取num的最后一位 res左移 和 num的最后一位进行 & 操作

注意：不能for num > 0 这样操作
同190题
*/
func reverseBits(num uint32) uint32 {
	var res uint32 = 0

	// 这种写法是错误的，若num的前几位为0，则直接退出了，无法保证得到的结果的位数与num的位数一致
	//for num > 0 {
	//	res = (res << 1) | (num & 1)
	//	num = num >> 1
	//}
	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num = num >> 1
	}
	return res
}
func main() {
	var num uint32 = 00011010
	fmt.Println(num)
	fmt.Println(reverseBits(num))
}
