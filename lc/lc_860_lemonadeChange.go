package main

import "fmt"

/*
一个糖果值5元钱，用户只有面值为5、10、20的钞票，一次只能买一个糖果
所以用户来你这买东西是否可以，即买到糖果并且得到相应的找零
开始你手中是没有钱的

第一种方法：用数组have表示你手中5、10的个数。根据用户给的钱来对5 10的个数进行调整，若面值为5，则直接对have[0]++
若面值为10，则判断have[0]的个数是否大于0
若面值为20，则判断have[0]和have[1]的个数是否都不为0或者have[0]的个数是否大于2
*/
func lemonadeChange(bills []int) bool {
	if len(bills) < 1 {
		return true
	}
	if bills[0] != 5 {
		return false
	}
	have := []int{1, 0}
	for i := 1; i < len(bills); i++ {
		if bills[i]==5 {
			have[0]++
		} else if bills[i]==10 {
			if have[0] > 0 {
				have[0]--
				have[1]++
			} else {
				return false
			}
		} else {
			if have[0] > 0 && have[1] > 0 {
				have[0]--
				have[1]--
			} else if have[0] > 2 {
				have[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
/*
思路同上面，每次多更新five、ten的个数，然后对five的个数判断是否小于0，若小于则返回false
*/
func lemonadeChange1(bills []int) bool {
	five, ten := 0, 0
	if len(bills) < 1 {
		return true
	}
	if bills[0] != 5 {
		return false
	}

	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			ten++
			five--
		} else if ten > 0 {
			ten--
			five--
		} else {
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func main() {
	bills := []int{5, 5, 5, 20, 20, 5, 5, 5}
	fmt.Println(lemonadeChange(bills), lemonadeChange1(bills))
}
