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
	var i, l int
	for _, childNode := range n.children {
		l = elementsFound + childNode.parentEdge.length
		edgeLabel := t.edgeValues[childNode.parentEdge.id][elementsFound:l]

		for ; i < len(key)-elementsFound; i++ {
			if key[i] != edgeLabel[i] {
				break
			}
		}

		if i != 0 {
			return childNode, i
		}
	}
	return n, 0
}

func (t *PatriciaTrie) search(key string) (*node, int, int) {
	currentNode := t.root
	elementsFound := 0
	lenKey := len(key)

	var i int
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
			currentNode, i = t.findPrefix(currentNode, key, elementsFound)
			elementsFound += i
			return currentNode, elementsFound, i
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

	currentNode, elementsFound, i := t.search(key)
	if currentNode == nil {
		currentNode = t.root
	}

	remainder := lenKey - elementsFound
	if elementsFound == 0 && i == 0 {
		t.edgeValues = append(t.edgeValues, key)
		edge := &edge{id: len(t.edgeValues) - 1, length: len(key)}
		childNode := &node{parentEdge: edge}
		currentNode.children = append(currentNode.children, childNode)
	} else if remainder > 0 {
		idx := currentNode.parentEdge.id
		if i != 0 {
			edgeLabel := t.edgeValues[currentNode.parentEdge.id]
			t.edgeValues = append(t.edgeValues, edgeLabel)
			edge := &edge{id: idx, length: currentNode.parentEdge.length - i}
			childNode := &node{parentEdge: edge}
			childNode.children = currentNode.children
			currentNode.children = []*node{childNode}
			currentNode.parentEdge.length = i
			idx = len(t.edgeValues) - 1
		}

		if currentNode.isLeaf() || i != 0 {
			t.edgeValues[idx] = key
			edge := &edge{id: idx, length: lenKey - elementsFound}
			childNode := &node{parentEdge: edge}
			currentNode.children = append(currentNode.children, childNode)
		} else {
			t.edgeValues = append(t.edgeValues, key)
			idx = len(t.edgeValues) - 1
			edge := &edge{id: idx, length: lenKey - elementsFound}
			childNode := &node{parentEdge: edge}
			currentNode.children = append(currentNode.children, childNode)
		}
	}
}

func (t *PatriciaTrie) Search(key string) bool {
	key += string('\x00')
	currentNode := t.root
	elementsFound := 0
	lenKey := len(key)

	var nextNode *node
	for currentNode != nil {
		if elementsFound == lenKey {
			break
		}

		if currentNode.children == nil {
			break
		}

		nextNode = nil
		for _, childNode := range currentNode.children {
			edgeLabel := t.edgeValues[childNode.parentEdge.id][elementsFound : elementsFound+childNode.parentEdge.length]

			if strings.HasPrefix(key, edgeLabel) {
				nextNode = childNode
				elementsFound += nextNode.parentEdge.length
				key = key[nextNode.parentEdge.length:]
				break
			}
		}
		currentNode = nextNode
	}

	return elementsFound == lenKey
}
