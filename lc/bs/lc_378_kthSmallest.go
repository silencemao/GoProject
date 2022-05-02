package main

/*
有矩阵第k小的元素
给你一个n x n矩阵matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。

要注意的地方是mid=l+(r-l)>>1，mid不一定是矩阵中的元素
那结果如何保证返回的一定是矩阵中的元素呢？
首先，我们要保证l、r中至少有一个元素是矩阵中的元素
其次，退出循环的条件是l==r，结合第一点，则返回值一定是矩阵中的元素

所以我们在计算过程中，当cnt==k的时候，并不是直接返回，而是继续寻找，直到l==r
当cnt<k时，表示<=mid的数字个数不到k个，因此l可以从mid后一个位置开始，此时r=r(原矩阵中的数字)
当cnt>=k时，表示<=mid的数字个数至少是k个，所以r从mid开始向前找，此时l=l(原矩阵中的数字), r=mid
*/
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		mid := l + (r-l)>>1
		cnt := getCount(mid, matrix)
		if cnt < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func getCount(num int, matrix [][]int) int {
	cnt, n := 0, len(matrix)
	for j := 0; j < n; j++ {
		for i := n - 1; i >= 0; i++ {
			if matrix[i][j] <= num {
				cnt += i + 1
				break
			}
		}
	}
	return cnt
}

func main() {

}
