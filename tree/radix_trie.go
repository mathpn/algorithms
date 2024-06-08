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
	label  int
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
		edgeLabel := t.edgeValues[currentNode.parentEdge.label][length : length+currentNode.parentEdge.length]
		edgeLabel = strings.Replace(edgeLabel, string('\x00'), "$", 1)
		path = append(path, fmt.Sprintf("[%d] %s", currentNode.parentEdge.label, edgeLabel))
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

func (t *PatriciaTrie) Insert(key string) {
	key += string('\x00')
	currentNode := t.root
	elementsFound := 0
	lenKey := len(key)
	fullKey := key // XXX

	var nextNode *node
	i := 0
	for currentNode != nil {
		if elementsFound == lenKey {
			break
		}

		if currentNode.children == nil {
			break
		}

		i = 0
		nextNode = nil
		for _, childNode := range currentNode.children {
			edgeLabel := t.edgeValues[childNode.parentEdge.label][elementsFound : elementsFound+childNode.parentEdge.length]

			if strings.HasPrefix(key, edgeLabel) {
				nextNode = childNode
				elementsFound += nextNode.parentEdge.length
				key = key[nextNode.parentEdge.length:]
				break
			}

			for ; i < lenKey-elementsFound; i++ {
				if key[i] != edgeLabel[i] {
					break
				}
			}
			if i != 0 {
				nextNode = childNode
				elementsFound += i
				break
			}
		}
		if nextNode == nil {
			break
		}
		currentNode = nextNode
		if i != 0 {
			break
		}
	}

	if currentNode == nil {
		currentNode = t.root
	}

	remainder := lenKey - elementsFound
	if elementsFound == 0 && i == 0 {
		t.edgeValues = append(t.edgeValues, key)
		edge := &edge{label: len(t.edgeValues) - 1, length: len(key)}
		childNode := &node{parentEdge: edge}
		currentNode.children = append(currentNode.children, childNode)
	} else if remainder > 0 {
		idx := currentNode.parentEdge.label
		if i != 0 {
			edgeLabel := t.edgeValues[currentNode.parentEdge.label]
			t.edgeValues = append(t.edgeValues, edgeLabel)
			edge := &edge{label: idx, length: currentNode.parentEdge.length - i}
			childNode := &node{parentEdge: edge}
			childNode.children = currentNode.children
			currentNode.children = []*node{childNode}
			currentNode.parentEdge.length = i
			idx = len(t.edgeValues) - 1
		}

		if currentNode.isLeaf() || i != 0 {
			t.edgeValues[idx] = fullKey
			edge := &edge{label: idx, length: len(key) - i}
			childNode := &node{parentEdge: edge}
			currentNode.children = append(currentNode.children, childNode)
		} else {
			t.edgeValues = append(t.edgeValues, fullKey)
			idx = len(t.edgeValues) - 1
			edge := &edge{label: idx, length: len(key) - i}
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
			edgeLabel := t.edgeValues[childNode.parentEdge.label][elementsFound : elementsFound+childNode.parentEdge.length]

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
