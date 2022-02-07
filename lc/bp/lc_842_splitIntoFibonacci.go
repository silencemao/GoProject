package main

import (
	"fmt"
	"strconv"
)

/*
给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。

形式上，斐波那契式序列是一个非负整数列表 F，且满足：

0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
F.length >= 3；
对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。

返回从 S 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。

示例 1：
输入："123456579"
输出：[123,456,579]

示例 2：
输入: "11235813"
输出: [1,1,2,3,5,8,13]

1、以ind为起点，[ind:ind+1]、[ind:i+2]之间的子串分别构成不同的数字 tmpNum ，判断每一个数字是否能够加入到结果中
2、当res中数字<=2时，或tmpNum和res[-1]、res[-2]构成斐波那契数列，直接将tmpNum加入到res中
3、当tmpNum >= 1<<31-1、或tmpNum>res[-1]+res[-2]时 直接结束当前循环
4、
*/
func help842(num string, ind int, res *[]int) bool {
	fmt.Println(res)
	if ind >= len(num) && len(*res) >= 3 {
		return true
	}
	for i := ind; i < len(num); i++ {
		if num[ind] == '0' && i > ind {
			break
		}

		tmpNum := substring(num, ind, i+1)
		if tmpNum >= 1<<31-1 {
			break
		}

		size := len(*res)
		if size > 2 && tmpNum > (*res)[size-1]+(*res)[size-2] {
			break
		}

		if len(*res) < 2 || tmpNum == (*res)[size-1]+(*res)[size-2] {
			*res = append(*res, tmpNum)
			if help842(num, i+1, res) {
				return true
			}
			*res = (*res)[:len(*res)-1]
		}
	}
	return false
}

func substring(num string, s, e int) int {
	res := num[s:e]
	r, _ := strconv.Atoi(res)
	return r
}

func splitIntoFibonacci(num string) []int {
	var res []int
	help842(num, 0, &res)
	return res
}

func main() {
	num := "74912134825162255812723932620170946950766784234934"
	fmt.Println(splitIntoFibonacci(num))
}
