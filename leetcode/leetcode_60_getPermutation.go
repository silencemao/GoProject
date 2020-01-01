package main

import (
	"fmt"
	"strconv"
)
// 获得第k个全排列
func getPermuteHelper(res *string, used *[]bool, p []int, n int, k *int) {
	if len(p)==n {
		*k--
		if *k==0 {
			for _, num := range p {
				*res += strconv.Itoa(num)
			}
		}
		return
	}
	for i := 1; i <= n; i++ {
		if !(*used)[i-1] {
			(*used)[i-1] = true
			p = append(p, i)
			getPermuteHelper(res, used, p, n, k)
			(*used)[i-1] = false
			p = p[:len(p)-1]
		}
	}
}
func getPermutation(n, k int) string {
	var p []int
	var res string
	var used  = make([]bool, n)
	getPermuteHelper(&res, &used, p, n, &k)

	return res
}

/*
n个数字，最高位确定之后，后面的数字有(n-1)!种组合，此高位确定之后，后面的数字有(n-2)种组合
所以我们要从最高位开始，逐个的确定每一个位置的数字
如何确定最高位的数字，以n=4为例，最高位取1，后面3位有6种组合，最高位取2后面3位有6种组合...
我们就要看k对6取整之后的数字，相当于k包含多少个 6种组合，取整之后的值就是我们要确定的最高位的数字索引
然后把选定的数字从候选集中删除。
接着对k取余，然后获得次高位的数字，以此类推
*/
func getPermutation1(n, k int) string{
	var nums  = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var factor []int
	var res string
	k -= 1         // k是从1开始计数的
	factor = append(factor, 1)
	for i := 1; i < n; i++ {
		factor = append(factor, factor[len(factor) - 1] * i)
	}
	for i := n; i >= 1; i-- {
		j := k / factor[i - 1]
		res += nums[j]
		k = k % factor[i - 1]
		nums = append(nums[:j], nums[j+1:]...)
	}
	return res
}

func main() {
	fmt.Println(getPermutation(4, 6))
	fmt.Println(getPermutation1(4, 6))
}