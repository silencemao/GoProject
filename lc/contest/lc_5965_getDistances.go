package main

import "fmt"

func abs(x int) int64 {
	if x < 0 {
		return int64(-x)
	}
	return int64(x)
}

func getDistances(arr []int) []int64 {
	res := make(map[int][]int, len(arr))
	answer := make([]int64, len(arr))
	for i, v := range arr {
		if preInds, ok := res[v]; ok {
			if len(preInds) == 1 {
				ind := preInds[0]
				answer[ind] += abs(ind - i)
				answer[i] += abs(ind - i)
			} else {
				for _, ind := range preInds {
					answer[ind] += abs(ind - i)
					answer[i] += abs(ind - i)
				}
			}
			preInds = append(preInds, i)
			res[v] = preInds
		} else {
			res[v] = []int{i}
		}
	}
	return answer
}

func getDistances1(arr []int) []int64 {
	res := make(map[int][]int, len(arr))
	answer := make([]int64, len(arr))
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		if preInd, ok := res[arr[i]]; ok {
			for _, ind := range preInd {
				answer[i] += abs(i - ind)
				answer[ind] += abs(i - ind)
			}
			preInd = append(preInd, i)
			res[arr[i]] = preInd
		} else {
			res[arr[i]] = []int{i}
		}

		if preInd, ok := res[arr[j]]; ok && i != j {
			for _, ind := range preInd {
				answer[j] += abs(j - ind)
				answer[ind] += abs(j - ind)
			}
			preInd = append(preInd, j)
			res[arr[j]] = preInd
		} else {
			res[arr[j]] = []int{j}
		}
	}
	return answer
}

func getDistances2(arr []int) []int64 {
	answer := make([]int64, len(arr))
	for i, v1 := range arr {
		for j := i + 1; j < len(arr); j += 1 {
			v2 := arr[j]
			if v1 == v2 {
				answer[i] += abs(i - j)
				answer[j] += abs(i - j)
			}
		}
	}
	return answer
}

// 前缀和+后缀和
func getDistances3(arr []int) []int64 {
	pre := make(map[int][]int, 0) //key为数组中的值，value表示上一个值为key的元素的下标 和 当前位置前面值为key的元素数量
	m1 := make([]int, len(arr))

	post := make(map[int][]int, 0)
	m2 := make([]int, len(arr))

	answer := make([]int64, len(arr))
	for i, v := range arr {
		if tmp, ok := pre[v]; ok {
			preInd, cnt := tmp[0], tmp[1]        // 索引i前一个值为v的元素的索引，在索引i之前值为v的元素数量，
			m1[i] += (i-preInd)*cnt + m1[preInd] //

			tmp[0] = i
			tmp[1] += 1
		} else {
			tmp = []int{i, 1}
			pre[v] = tmp
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		if tmp, ok := post[v]; ok {
			postInd, cnt := tmp[0], tmp[1] // 索引i之后值为v的元素索引，索引i之后只为v的元素数量
			m2[i] += (postInd-i)*cnt + m2[postInd]

			tmp[0] = i
			tmp[1] += 1
		} else {
			tmp = []int{i, 1}
			post[v] = tmp
		}
	}
	for i := 0; i < len(arr); i++ {
		answer[i] = int64(m1[i] + m2[i])
	}
	return answer
}

func main() {
	arr := []int{2, 1, 3, 1, 2, 3, 6, 3, 1, 3}
	fmt.Println(getDistances(arr))
	fmt.Println(getDistances1(arr))
	fmt.Println(getDistances2(arr))
	fmt.Println(getDistances3(arr))
}
