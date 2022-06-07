package _struct

/*
前缀树，字典树
*/
type Trie struct {
	isEnd bool
	next  []*Trie
}

func Constructor() Trie {
	trie := Trie{}
	trie.isEnd = false
	trie.next = make([]*Trie, 26)
	return trie
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		ch := c - 'a'
		if node.next[ch] == nil {
			t := Constructor()
			node.next[ch] = &t
		}
		node = node.next[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		ch := c - 'a'
		node = node.next[ch]
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		ch := c - 'a'
		node = node.next[ch]
		if node == nil {
			return false
		}
	}
	return true
}
