package main

import (
	"fmt"
	"math"
)

func consecutiveNumbersSum(N int) int {
	res := 1
	for i := 2; i <= int(math.Sqrt(float64(2*N))); i++ {
		if (N - i * (i - 1) / 2) % i == 0 {
			res++
		}
	}
	return res
}

//func main() {
//	fmt.Print(consecutiveNumbersSum(9))
//}
