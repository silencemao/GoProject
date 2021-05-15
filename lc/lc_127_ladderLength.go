package main

import "fmt"

/*

给定一个字符串beginWord、endWord 以及字符串列表wordList中
从beginWord开始，每次变换一个字母，并且新生成的字母在wordList中，算作一次变换，需要多少次变换才能将beginWord变换成endWord

利用一个队列queue记录beginWord - endWord的变化过程
首先将beginWord放入queue中，
每次取queue的队首作为beginWord 然后尝试将beginWord中的每一个字母从a-z进行变化，若变化之后等于endWord，则直接return
若不等于endWord，但是在wordList中并且之前没有使用过，表示当前的word可以作为一个过度，然后放到queue中，
下次再尝试将beginWord继续调整，直到变化为endWord

*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet, pathCnt := make(map[string]bool), make(map[string]int)

	pathCnt[beginWord] = 1
	queue := []string{beginWord}

	for _, word := range wordList {
		wordSet[word] = true
	}

	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]

		for i := 0; i < len(word); i++ {
			newWord := []byte(word)
			for j := 'a'; j <= 'z'; j++ {
				if byte(j) == newWord[i] {
					continue
				}
				newWord[i] = byte(j)
				newWordStr := string(newWord)

				if wordSet[newWordStr] && newWordStr == endWord {
					return pathCnt[word] + 1
				}

				if _, ok := pathCnt[newWordStr]; !ok && wordSet[newWordStr] {
					queue = append(queue, newWordStr)
					pathCnt[newWordStr] = pathCnt[word] + 1
				}
			}
		}
	}
	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))

}
