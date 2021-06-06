package main

import "fmt"

/*
给定一个数组，数组中只有两个数字只出现了一次，其余数字均出现了两次
在O(n)的时间内找出这两个数字

前提知识：
1、异或 相同为0不同为1
2、一个数和其负数进行与&操作，肯定是2的指数次幂(2^n)，即二进制结果中只有一个1，这个1肯定是在原来数字和在负数中都为1

假设只出现一次的数字为a和b
1、数组中只有a、b只出现了一次，其余数字数字出现了两次。首先将数组中的数字进行异或，得到的结果即为a^b之后的结果diff
2、然后我们将diff与-diff进行&操作。diff &=-diff 此时diff的二进制表达式中只有一个位为1
3、然后将diff与数组中的数字再去进行异或操作，这时肯定有一部分数字和diff异或之后为0，另外一部分不为0
4、这样就能将数字分成两部分，其中一部分包含a和部分的出现两次的数字，另外一部分包含b和部分出现两次的数字
5、出现两次的数字异或之后还是0，则只剩下a和b了

注意：在diff第二次与num异或时，结果分为0和非0部分，非0部分不一定是1，所以在判断的时候要注意

*/
func singleNumber260(nums []int) []int {
	res := []int{0, 0}
	diff := 0
	for _, num := range nums {
		diff ^= num
	}

	diff = diff & (-diff)
	for _, num := range nums {
		//fmt.Println(diff&num, num)
		if diff&num == 0 {
			res[0] ^= num
		} else { // 此处不一定为1
			res[1] ^= num
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber260(nums))
}
