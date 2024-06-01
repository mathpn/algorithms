package tree

type trieNode struct {
	childNode []*trieNode
	wordCount int
}

func NewTrieNode() *trieNode {
	return &trieNode{childNode: make([]*trieNode, 26), wordCount: 0}
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(key string) {
	node := t.root
	for _, char := range key {
		childNode := node.childNode[char-'a']
		if childNode == nil {
			childNode = NewTrieNode()
			node.childNode[char-'a'] = childNode
		}
		node = node.childNode[char-'a']
	}
	node.wordCount++
}

func (t *Trie) Search(key string) bool {
	node := t.root
	for _, char := range key {
		childNode := node.childNode[char-'a']
		if childNode == nil {
			return false
		}
		node = childNode
	}
	return node.wordCount > 0
}

func (t *Trie) StartsWith(key string) bool {
	node := t.root
	for _, char := range key {
		childNode := node.childNode[char-'a']
		if childNode == nil {
			return false
		}
		node = childNode
	}
	return true
}
