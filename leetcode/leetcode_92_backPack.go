package main

import (
	"fmt"
	"myGoProject/leetcode/util"
)

func backPack(m int, A []int) int {
	if m <= 0 || len(A) < 1 {
		return 0
	}

	res := make([]int, m+1)
	for _, w := range A {
		for j := m; j > 0; j-- {
			if j >= w {
				//res[j] = int(math.Max(float64(res[j-w]+w), float64(res[j])))
				res[j] = util.MaxInt(res[j - w] + w, res[j])
			}
		}
	}
	return res[m]
}

func main() {
	A := []int{3 ,4, 8, 5}
	m := 10
	fmt.Println(backPack(m, A))
}
