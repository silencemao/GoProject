package main

/*
两个字符串类型的数字相乘
可以按照乘法从低位到高位逐渐相乘的思路来执行，而且两数相乘得到的数字长度是不会超过；两个数数字的长度之和
所以可以先定义一个长度为l1+l2的数组，然后分别从num1和num2的低位开始相乘，乘完之后我们需要将结果的个位和十位拆分开来，
分别放到i+j+1, i+j上。记得再拆分之前要先加上进位的数字
*/

func multiply(num1 string, num2 string) string {
	if num1=="0" || num2=="0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)
	var array = make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2-1; j >=0; j-- {
			p1, p2 := i + j, i + j + 1
			sum := (num1[i] - '0') * (num2[j] - '0') + array[p2]  // p2相较于p1是低位，所以加p2位置上的数
			array[p1] += sum / 10
			array[p2] = sum % 10
		}
	}

	var index int
	for index = 0; index < len(array)-1; index++ {
		if array[index] != 0 {
			break
		}
	}
	for i:= index; i < len(array); i++ {
		array[i] += '0'
	}
	return string(array[index:])
}

//func main() {
//	num1, num2 := "123", "456"
//	fmt.Println(multiply(num1, num2))
//}
