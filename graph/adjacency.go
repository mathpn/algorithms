package graph

import (
	"fmt"
	"main/list"

	"golang.org/x/exp/slices"
)

type AdjacencyMatrix[T comparable] [][]T

func BFS(matrix AdjacencyMatrix[int], source int, needle int) ([]int, error) {
	l := len(matrix)
	seen := make([]bool, l)
	prev := make([]int, l)
	path := []int{source}
	for i := range prev {
		prev[i] = -1
	}
	if source < 0 || source >= l {
		return path, fmt.Errorf("source %d is out of bounds of graph with %d nodes", source, l)
	}
	seen[source] = true
	q := list.Queue[int]{}
	q.Enqueue(source)
	var curr int
	var adjs []int
	var err error
	for q.Len > 0 {
		curr, err = q.Dequeue()
		if err != nil {
			return path, err
		}
		if curr == needle {
			break
		}
		adjs = matrix[curr]
		for i := 0; i < len(matrix); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}
			seen[i] = true
			prev[i] = curr
			q.Enqueue(i)
		}
		seen[curr] = true
	}

	if prev[needle] == -1 {
		return path, fmt.Errorf("needle %d not found in graph", needle)
	}

	curr = needle
	out := make([]int, 0)
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	slices.Reverse(out)
	return append(path, out...), nil
}

type GraphEdge struct {
	to     int
	weight float64
}

type AdjacencyList [][]GraphEdge

func walk(graph AdjacencyList, curr int, needle int, seen []bool, path *list.Stack[int]) (bool, error) {
	if seen[curr] {
		return false, nil
	}

	seen[curr] = true
	path.Push(curr)
	if curr == needle {
		return true, nil
	}

	list := graph[curr]
	var edge GraphEdge
	var found bool
	var err error
	for i := 0; i < len(list); i++ {
		edge = list[i]
		found, err = walk(graph, edge.to, needle, seen, path)
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
	}

	_, err = path.Pop()
	if err != nil {
		return false, err
	}
	return false, nil
}

func DFS(graph AdjacencyList, source int, needle int) ([]int, error) {
	seen := make([]bool, len(graph))
	path := list.Stack[int]{}
	outPath := make([]int, 0)
	found, err := walk(graph, source, needle, seen, &path)
	if err != nil {
		return nil, err
	}
	if !found {
		return outPath, fmt.Errorf("value %d not found in graph", needle)
	}
	var v int
	for path.Len > 0 {
		v, err = path.Pop()
		if err != nil {
			return outPath, err
		}
		outPath = append(outPath, v)
	}
	slices.Reverse(outPath)
	return outPath, nil
}
