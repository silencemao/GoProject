package main

import (
	"fmt"
	"strconv"
)

var cnt int

func helper(pStr string, pInd int) {
	if pInd > len(pStr) {
		return
	}
	if pInd==len(pStr) {
		fmt.Println(cnt, pStr)
		cnt++
	}
	tStr := []byte(pStr)
	for i := pInd; i < len(tStr); i++ {
		tStr[i], tStr[pInd] = tStr[pInd], tStr[i]
		helper(string(tStr), pInd+1)
		tStr[i], tStr[pInd] = tStr[pInd], tStr[i]
	}
}

func permutation(n int) {
	tStr := ""
	for i := 1; i <= n; i++ {
		tStr += strconv.Itoa(i)
	}
	helper(tStr, 0)
}

func main() {
	permutation(4)
}
