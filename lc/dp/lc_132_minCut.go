package main

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
		for j := range g[i] {
			g[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
		}
	}

	f := make([]int, n)
	for i := range f {
		if g[0][i] { // 0-i之间子串已经是回文了，不需要切割
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if g[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	return f[n-1]
}

func main() {
	s := "aab"
	fmt.Println(minCut(s))
}
