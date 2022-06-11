package _struct

/*
前缀树，字典树
*/
type Trie struct {
	IsEnd bool
	Next  []*Trie
}

func Constructor() Trie {
	trie := Trie{}
	trie.IsEnd = false
	trie.Next = make([]*Trie, 26)
	return trie
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		ch := c - 'a'
		if node.Next[ch] == nil {
			t := Constructor()
			node.Next[ch] = &t
		}
		node = node.Next[ch]
	}
	node.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		ch := c - 'a'
		node = node.Next[ch]
		if node == nil {
			return false
		}
	}
	return node.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		ch := c - 'a'
		node = node.Next[ch]
		if node == nil {
			return false
		}
	}
	return true
}
