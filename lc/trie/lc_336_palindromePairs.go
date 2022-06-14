package main

import "fmt"

/*
给定一组 互不相同 的单词， 找出所有 不同的索引对 (i, j)，使得列表中的两个单词，words[i] + words[j]，可拼接成回文串。

输入：words = ["abcd","dcba","lls","s","sssll"]
输出：[[0,1],[1,0],[3,2],[2,4]]
解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
*/
/*
1、暴力解法
2、hash表
3、字典树
*/
func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

//超时
func palindromePairs(words []string) [][]int {
	var res [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if isPalindrome(words[i] + words[j]) {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func findWord(word string, set map[string]int) int {
	if ind, ok := set[word]; ok {
		return ind
	}
	return -1
}

func palindromePairs1(words []string) [][]int {
	set := make(map[string]int, 0)
	for i, word := range words {
		a := []rune(word)
		for m, n := 0, len(a)-1; m < n; m, n = m+1, n-1 {
			a[m], a[n] = a[n], a[m]
		}
		set[string(a)] = i
	}
	var res [][]int
	for i := 0; i < len(words); i++ {
		word := words[i]
		m := len(word)

		for j := 0; j <= m; j++ {
			if isPalindrome(word[j:m]) {
				leftId := findWord(word[:j], set)
				if leftId != -1 && leftId != i {
					res = append(res, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word[:j]) {
				rightId := findWord(word[j:], set)
				if rightId != -1 && rightId != i {
					res = append(res, []int{rightId, i})
				}
			}
		}
	}
	return res
}

// 字典树
type Node struct {
	Ind  int
	Next []*Node
}

func (n *Node) insert(word string, i int) {
	for _, c := range word {
		if n.Next[c-'a'] == nil {
			n1 := new(Node)
			n1.Ind = -1
			n1.Next = make([]*Node, 26)
			n.Next[c-'a'] = n1
		}
		n = n.Next[c-'a']
	}
	n.Ind = i
}

func (n *Node) find(word string) int {
	for _, c := range word {
		if n.Next[c-'a'] == nil {
			return -1
		}
		n = n.Next[c-'a']
	}
	return n.Ind
}

func palindromePairs2(words []string) [][]int {
	node := new(Node)
	node.Next = make([]*Node, 26)
	node.Ind = -1
	for i, word := range words {
		a := []rune(word)
		for m, n := 0, len(a)-1; m < n; m, n = m+1, n-1 {
			a[m], a[n] = a[n], a[m]
		}
		node.insert(string(a), i)
	}
	var res [][]int
	for i := 0; i < len(words); i++ {
		word := words[i]
		m := len(word)

		for j := 0; j <= m; j++ {
			if isPalindrome(word[j:m]) {
				leftId := node.find(word[:j])
				if leftId != -1 && leftId != i {
					res = append(res, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word[:j]) {
				rightId := node.find(word[j:])
				if rightId != -1 && rightId != i {
					res = append(res, []int{rightId, i})
				}
			}
		}
	}
	return res
}

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words))
	fmt.Println(palindromePairs1(words))
	fmt.Println(palindromePairs2(words))
}
