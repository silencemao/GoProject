package main

import "fmt"

//统计所有小于非负整数 n 的质数的数量。
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func countPrimes1(n int) int {
	isprime := make([]bool, n)
	for i := 0; i < n; i++ {
		isprime[i] = true
	}
	for i := 2; i < n; i++ {
		if isprime[i] {
			for j := 2; j*i < n; j++ {
				isprime[j*i] = false
			}
		}
	}
	res := 0
	for i := 2; i < n; i++ {
		if isprime[i] {
			res++
		}
	}
	return res
}

func main() {
	n := 2
	fmt.Println(countPrimes(n))
	fmt.Println(countPrimes1(n))
}
