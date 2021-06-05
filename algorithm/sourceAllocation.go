package main

import (
	"fmt"
	"math"
	"math/rand"
)

var r = rand.New(rand.NewSource(0)) // 用于搜索

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 计算score
func calculateScore(nums []int) int {
	sum, res := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for j := i + 1; j < len(nums); j++ {
			res += abs(nums[j] - nums[i])
		}
	}
	sum /= len(nums)
	for i := 0; i < len(nums); i++ {
		res += abs(sum - nums[i])
	}
	return res
}

func dispatch(nums []int, m int) {
	T, minT, ratio := 100.0, 0.0001, 0.99 // 初始化最高温度，最低温度，降温速率

	tmp := append([]int{}, nums...)
	fix := initial(len(nums), m) // 将m分配完成
	for i := 0; i < len(nums); i++ {
		tmp[i] += fix[i]
	}
	s1 := calculateScore(tmp)          // 计算score
	bestNum := append([]int{}, tmp...) // 存储最优解
	fmt.Println(bestNum)

	for T > minT {
		iter := 500
		for iter > 0 {
			pre := append([]int{}, tmp...)

			localSearch(nums, pre) // 搜索
			s2 := calculateScore(pre)
			if s2 < s1 { // 若新解优于最优解，直接接受
				bestNum = append([]int{}, pre...)
				tmp, s1 = append([]int{}, pre...), s2
			} else if math.Exp(-float64(s2-s1)/T) < r.Float64() { // 有概率接受
				tmp, s1 = append([]int{}, pre...), s2
			}
			iter--
		}
		T *= ratio
	}
	//nums = bestNum  不会生效
	for i := 0; i < len(nums); i++ {
		nums[i] = bestNum[i]
	}
	fmt.Println(bestNum)
}

// 初始化 参数
func initial(cnt, m int) []int {
	var res []int
	for i := 0; i < cnt-1; i++ {
		num := r.Intn(m)
		m -= num
		res = append(res, num)
	}
	res = append(res, m)
	return res
}

func findMaxMin(nums []int) (int, int) {
	min, max := 1<<32-1, -1<<32
	ind1, ind2 := 0, 0

	for ind, num := range nums {
		if num > max {
			max, ind2 = num, ind
		}
		if num < min {
			min, ind1 = num, ind
		}
	}
	return ind1, ind2
}

func localSearch(nums, tmp []int) {
	minInd, maxInd := findMaxMin(tmp)
	if nums[minInd] > tmp[minInd] || nums[maxInd] > tmp[maxInd] {
		tmp = append([]int{}, nums...)
		return
	}
	if tmp[maxInd] > nums[maxInd] { // 防止r.Intn() 报错
		m := r.Intn(tmp[maxInd] - nums[maxInd])
		tmp[minInd] += m
		tmp[maxInd] -= m
	}
}

func calculateSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	nums := []int{1, 200, 3, 4, 50}
	m := 200
	sum := m + calculateSum(nums)
	fmt.Println(nums, sum)

	dispatch(nums, m)
	fmt.Println(nums, calculateSum(nums))
}
