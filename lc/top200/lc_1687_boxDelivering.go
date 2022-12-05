package main

import "GoProject/leetcode/util"

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum := 0
		dp[i] = 1<<31 - 1
		for j := i; j >= 1 && j >= i-maxBoxes+1; j-- {
			sum += boxes[j-1][1]
			if sum > maxWeight {
				break
			}
			dp[i] = util.MinInt(dp[i], dp[j-1]+cost1687(boxes, j, i))
		}
	}
	return dp[n]
}

func cost1687(boxes [][]int, l, r int) int {
	ans, port := 2, boxes[l-1][0]
	for l+1 <= r {
		l += 1
		if boxes[l-1][0] != port {
			ans += 1
		}
		port = boxes[l-1][0]
	}
	return ans
}

/*
定义 f[i]表示将前i个箱子搬运到码头需要的趟次数，则最终结果是f[n]
枚举上一次运输的最后一个箱子的编号，则f[i]可以由f[j]转移而来 f[i] = min(f[i], f[j]+cost(j+1, i))
cost(j+1, i) = cost(i) - cost(j)
*/
func boxDelivering1(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	ws := make([]int, n+1) // 重量前缀和
	cs := make([]int, n)   // 趟次前缀和 cs[i]+2 表示配送完成前i个箱子需要的趟次数
	for i, box := range boxes {
		p, w := box[0], box[1]
		ws[i+1] = ws[i] + w
		if i < n-1 {
			t := 0
			if p != boxes[i+1][0] {
				t++
			}
			cs[i+1] = cs[i] + t
		}
	}
	f := make([]int, n+1)
	q := []int{0} // 构成f[q[i]] - cs[q[i]]的单调队列，存储完成前i个箱子需要的最少趟次数
	for i := 1; i <= n; i++ {
		for len(q) > 0 && (i-q[0] > maxBoxes || ws[i]-ws[q[0]] > maxWeight) {
			q = q[1:]
		}
		if len(q) > 0 {
			f[i] = cs[i-1] + f[q[0]] - cs[q[0]] + 2
		}
		if i < n {
			for len(q) > 0 && f[q[len(q)-1]]-cs[q[len(q)-1]] >= f[i]-cs[i] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
	}
	return f[n]
}

func main() {

}
