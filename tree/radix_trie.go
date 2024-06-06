package tree

import (
	"fmt"
	"strings"
)

type node struct {
	parentEdge *edge
	children   []*node
	isLeaf     bool
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

// FIXME this is WIP and completely wrong
func (t *PatriciaTrie) Insert(key string) {
	currentNode := t.root
	t.edgeValues = append(t.edgeValues, key)

	edge := &edge{label: len(t.edgeValues) - 1, length: len(key)}
	childNode := &node{parentEdge: edge, isLeaf: true}
	currentNode.children = append(currentNode.children, childNode)
}

func (t *PatriciaTrie) Search(key string) bool {
	currentNode := t.root
	elementsFound := 0
	lenKey := len(key)
	fullKey := key // XXX

	var nextNode *node
	for {
		fmt.Printf("-> %v | elementsFound: %d | key: %s\n", currentNode, elementsFound, key)
		if currentNode == nil {
			break
		}

		if elementsFound == lenKey {
			break
		}

		if currentNode.children == nil {
			break
		}

		nextNode = nil
		for _, childNode := range currentNode.children {
			fmt.Printf("key: %s\n", key)
			edgeLabel := t.edgeValues[childNode.parentEdge.label][elementsFound : elementsFound+childNode.parentEdge.length]
			fmt.Printf("edgeLabel: %s\n", edgeLabel)

			if strings.HasPrefix(key, edgeLabel) {
				nextNode = childNode
				elementsFound += nextNode.parentEdge.length
				key = key[nextNode.parentEdge.length:]
				break
			}
		}
		currentNode = nextNode
	}

	fmt.Printf("key: %s - elementsFound: %d | lenKey: %d\n", fullKey, elementsFound, lenKey)
	fmt.Println("-----------")
	return elementsFound == lenKey
}
