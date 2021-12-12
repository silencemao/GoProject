package main

import "fmt"

/*
https://leetcode-cn.com/problems/watering-plants-ii/
*/
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	a, b := capacityA, capacityB
	res := 0
	l, r := 0, len(plants)-1
	for l <= r {
		lastA := a
		if l < r || a >= b {
			if a >= plants[l] {
				a -= plants[l]
			} else {
				a = capacityA - plants[l]
				res += 1
			}
		}

		if l < r || lastA < b {
			if b >= plants[r] {
				b -= plants[r]
			} else {
				b = capacityB - plants[r]
				res += 1
			}
		}
		l, r = l+1, r-1
	}
	return res
}
func main() {
	plants := []int{2, 2, 5, 2, 2}
	capacityA, capacityB := 5, 5
	fmt.Println(minimumRefill(plants, capacityA, capacityB))
}
