package main

/*
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

*/
func isPerfectSquare(num int) bool {
	l, r := 1, num/2

	for l <= r {
		mid := l + (r-l)>>1
		tmp := mid * mid
		if tmp == num {
			return true
		}
		if tmp > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {

}
