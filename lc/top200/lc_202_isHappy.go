package main

import "fmt"

func help202(n int, m map[int]bool) bool {
	if n == 1 {
		return true
	}
	tmp := 0
	for n > 0 {
		tmp += (n % 10) * (n % 10)
		n /= 10
	}
	if m[tmp] {
		return false
	}
	m[tmp] = true
	return help202(tmp, m)
}

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果 可以变为 1，那么这个数就是快乐数。
如果 n 是快乐数就返回 true ；不是，则返回 false 。

利用hash表，存储每次变换经历的数字，防止出现死循环
*/
func isHappy(n int) bool {
	m := make(map[int]bool, 0)
	m[n] = true

	return help202(n, m)
}

// 还有一种做法 类似检测链表是否有环 快慢指针

func main() {
	n := 2
	fmt.Println(isHappy(n))
}
