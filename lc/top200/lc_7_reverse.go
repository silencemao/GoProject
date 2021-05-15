package main

import "fmt"

func reverse(x int) int {
	res := 0
	minValue, maxValue := -(1 << 31), (1<<31)-1

	for x != 0 {
		//if res > maxValue/10 || res < minValue/10 {
		//	return 0
		//}
		//res = res*10 + x%10
		//x = x / 10
		if res >= minValue/10 && res <= maxValue/10 {
			res = res*10 + x%10
			x = x / 10
		} else {
			return 0
		}
	}
	return res
}

func main() {
	a := 9463847412
	fmt.Println(reverse(a))
}
