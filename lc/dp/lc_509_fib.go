package main

/*
斐波那切数列
*/
func fib(n int) int {
	if n < 2 {
		return n
	}
	f, s := 0, 1
	for i := 2; i <= n; i++ {
		res := f + s
		f, s = s, res
	}
	return s
}

func main() {

}
