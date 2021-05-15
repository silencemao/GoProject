package main

import (
	"fmt"
)

/*

 */
func checkPerfectNumber(num int) bool {
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if num%i != i { // 每个被除数只能加和一次
				sum += num / i // 取另一个大于i的被除数
			}
		}
	}

	return num != 1 && sum == num
}

func main() {
	fmt.Println(checkPerfectNumber(28))
	fmt.Println(checkPerfectNumber(496))
	fmt.Println(checkPerfectNumber(8128))
	fmt.Println(checkPerfectNumber(33550336))
}
