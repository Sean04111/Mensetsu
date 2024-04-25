package structrue

// Trie 前缀树实现
type Trie struct {
	head *TrieNode
}

type TrieNode struct {
	val      int
	children [26]*TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.head
	for _, v := range word {
		ch := v - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				val:      -1,
				children: [26]*TrieNode{},
			}
		}
		node = node.children[ch]
	}
	node.val = 1
}

func (t *Trie) SearchWord(word string) bool {
	node := t.head
	for _, v := range word {
		ch := v - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.val == 1
}

func (t *Trie) SearchWithPrefix(prefix string) bool {
	node := t.head
	for _, v := range prefix {
		ch := v - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}
