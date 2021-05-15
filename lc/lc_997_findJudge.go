package main

import (
	"fmt"
)

/*
找出法官
有N个人，编号为1 2 3 .. N
其中一个是法官，法官不相信任何人，而任何人都相信法官，[a,b]表示a相信b
找出其中的法官，若不存在法官，返回-1
我们需要找出其中一个人，有N-1个人相信他，而他不相信任何人。那如何记录相信他的人的个数 以及他相信人的个数呢
我们可以利用一个数组记录第i个人相信其他人的个数
另一个数组记录别人相信第i个人的个数
*/
func findJudge(N int, trust [][]int) int {

	in, out := make([]int, N+1), make([]int, N+1)
	for _, arr := range trust {
		in[arr[1]]++
		out[arr[0]]++
	}
	for i := 0; i <= N; i++ {
		if in[i] == N-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	N := 4
	fmt.Println(findJudge(N, trust))
}
