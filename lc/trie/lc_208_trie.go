package main

import (
	_struct "GoProject/leetcode/struct"
	"fmt"
)

func t() {
	trie := _struct.Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))

	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}

func main() {
	t()
}
