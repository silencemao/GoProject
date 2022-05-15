package main

/*
给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干 素数 组成，且其中所有整数互不相同。

对于每对满足 0 <= i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。

那么第k个最小的分数是多少呢? 以长度为 2 的整数数组返回你的答案, 这里answer[0] == arr[i]且answer[1] == arr[j] 。

输入：arr = [1,2,3,5], k = 3
输出：[2,5]
解释：已构造好的分数,排序后如下所示:
1/5, 1/3, 2/5, 1/2, 3/5, 2/3
很明显第三个最小的分数是 2/5
*/

func kthSmallestPrimeFraction(arr []int, k int) []int {
	l, r, n := 0.0, 1.0, len(arr)
	for l < r {
		mid := (l + r) / 2
		x, y, cnt := 0, 1, 0
		for i := 0; i < n; i++ {
			j := i + 1
			for ; j < n; j++ {
				if float64(arr[i])/float64(arr[j]) <= mid {
					cnt += n - j
					break
				}
			}
			if arr[i]*y > arr[j]*x { // 维护以arr[i]为分子的最接近mid的分数
				x, y = arr[i], arr[j]
			}
		}
		if cnt == k {
			return []int{x, y}
		} else if cnt < k { //l r 的变化不是mid+1或mid-1了
			l = mid
		} else {
			r = mid
		}
	}
	return []int{-1, -1}
}

func main() {

}
