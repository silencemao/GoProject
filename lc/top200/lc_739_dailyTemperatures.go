package main

/*
每日温度
给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指对于第 i 天，
下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。

输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]
单调栈
*/
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var q []int
	q = append(q, 0)
	for i := 1; i < n; i++ {
		for len(q) > 0 && temperatures[i] > temperatures[q[len(q)-1]] {
			preInd := q[len(q)-1]
			ans[preInd] = i - preInd
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	return ans
}

func main() {

}
