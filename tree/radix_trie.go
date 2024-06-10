package tree

import (
	"fmt"
	"strings"
)

type node struct {
	parentEdge *edge
	children   []*node
}

type edge struct {
	id     int
	length int
}

type PatriciaTrie struct {
	root       *node
	edgeValues []string
}

func NewPatriciaTrie() *PatriciaTrie {
	return &PatriciaTrie{root: &node{}, edgeValues: make([]string, 0)}
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (t *PatriciaTrie) Print() {
	fmt.Println("-> TRIE:")
	fmt.Printf("%v\n", t.edgeValues)
	node := t.root
	t.print(node, 0, make([]string, 0))
}

func (t *PatriciaTrie) print(currentNode *node, length int, path []string) {
	if currentNode == nil {
		return
	}

	if currentNode.parentEdge != nil {
		// fmt.Println(t.edgeValues[currentNode.parentEdge.id])
		// fmt.Printf("-> idx %d -- %d : %d - length %d\n", currentNode.parentEdge.id, length, length+currentNode.parentEdge.length, currentNode.parentEdge.length)
		edgeLabel := t.edgeValues[currentNode.parentEdge.id][length : length+currentNode.parentEdge.length]
		edgeLabel = strings.Replace(edgeLabel, string('\x00'), "$", 1)
		path = append(path, fmt.Sprintf("[%d] %s", currentNode.parentEdge.id, edgeLabel))
		length += currentNode.parentEdge.length
	}

	if currentNode.isLeaf() {
		path := strings.Join(path, " -> ")
		fmt.Printf("PATH: %s\n", path)
		return
	}

	for _, childNode := range currentNode.children {
		t.print(childNode, length, path)
	}

}

func (t *PatriciaTrie) findChild(n *node, key string, elementsFound int) *node {
	var l int
	for _, childNode := range n.children {
		l = elementsFound + childNode.parentEdge.length
		edgeLabel := t.edgeValues[childNode.parentEdge.id][elementsFound:l]

		if strings.HasPrefix(key, edgeLabel) {
			return childNode
		}
	}
	return nil
}

func (t *PatriciaTrie) findPrefix(n *node, key string, elementsFound int) (*node, int) {
	var overlap, l int
	for _, childNode := range n.children {
		l = elementsFound + childNode.parentEdge.length
		edgeLabel := t.edgeValues[childNode.parentEdge.id][elementsFound:l]

		for ; overlap < len(key)-elementsFound; overlap++ {
			if key[overlap] != edgeLabel[overlap] {
				break
			}
		}

		if overlap != 0 {
			return childNode, overlap
		}
	}
	return n, 0
}

func (t *PatriciaTrie) search(key string) (*node, int, int) {
	currentNode := t.root
	elementsFound := 0
	lenKey := len(key)

	var overlap int
	var nextNode *node
	for currentNode != nil {
		if elementsFound == lenKey {
			break
		}

		if currentNode.children == nil {
			break
		}

		nextNode = nil
		nextNode = t.findChild(currentNode, key, elementsFound)
		if nextNode == nil {
			currentNode, overlap = t.findPrefix(currentNode, key, elementsFound)
			elementsFound += overlap
			return currentNode, elementsFound, overlap
		}
		key = key[nextNode.parentEdge.length:]
		elementsFound += nextNode.parentEdge.length
		currentNode = nextNode
	}

	return currentNode, elementsFound, 0
}

func (t *PatriciaTrie) Insert(key string) {
	key += string('\x00')
	lenKey := len(key)

	currentNode, elementsFound, overlap := t.search(key)
	if currentNode == nil {
		currentNode = t.root
	}

	if elementsFound == lenKey {
		return
	}

	if elementsFound == 0 {
		t.insertRootChild(currentNode, key)
	} else {
		t.insertNode(currentNode, key, elementsFound, overlap)
	}
}

func (t *PatriciaTrie) insertRootChild(n *node, key string) {
	t.edgeValues = append(t.edgeValues, key)
	edge := &edge{id: len(t.edgeValues) - 1, length: len(key)}
	childNode := &node{parentEdge: edge}
	n.children = append(n.children, childNode)
}

func (t *PatriciaTrie) insertNode(n *node, key string, elementsFound int, overlap int) {
	idx := n.parentEdge.id
	lenKey := len(key)

	if overlap != 0 {
		splitEdge := &edge{id: idx, length: n.parentEdge.length - overlap}
		splitNode := &node{parentEdge: splitEdge}
		splitNode.children = n.children
		n.children = []*node{splitNode}
		n.parentEdge.length = overlap
	}

	t.edgeValues = append(t.edgeValues, key)
	idx = len(t.edgeValues) - 1
	newEdge := &edge{id: idx, length: lenKey - elementsFound}
	newNode := &node{parentEdge: newEdge}
	n.children = append(n.children, newNode)
}

func (t *PatriciaTrie) Search(key string) bool {
	key += string('\x00')
	_, elementsFound, _ := t.search(key)
	return elementsFound == len(key)
}

func (t *PatriciaTrie) StartsWith(key string) bool {
	_, elementsFound, _ := t.search(key)
	return elementsFound == len(key)
}
