package graph

import "testing"

func createMatrix() AdjacencyMatrix[int] {
	return [][]int{
		{0, 3, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestBFSAdjacencyMatrix(t *testing.T) {
	matrix := createMatrix()

	path, err := BFS(matrix, 0, 6)
	if err != nil {
		t.Fatal(err)
	}
	exp := []int{0, 1, 4, 5, 6}
	if !slicesEqual(path, exp) {
		t.Errorf("expected path %v, received %v", exp, path)
	}
}

func createList() AdjacencyList {
	return AdjacencyList{
		[]GraphEdge{
			{to: 1, weight: 3},
			{to: 2, weight: 1},
		},
		[]GraphEdge{
			{to: 4, weight: 1},
		},
		[]GraphEdge{
			{to: 3, weight: 7},
		},
		[]GraphEdge{},
		[]GraphEdge{
			{to: 1, weight: 1},
			{to: 3, weight: 5},
			{to: 5, weight: 2},
		},
		[]GraphEdge{
			{to: 2, weight: 18},
			{to: 6, weight: 1},
		},
		[]GraphEdge{
			{to: 3, weight: 1},
		},
	}
}

func TestDFSAdjacencyList(t *testing.T) {
	list := createList()
	path, err := DFS(list, 0, 6)
	if err != nil {
		t.Fatal(err)
	}
	exp := []int{0, 1, 4, 5, 6}
	if !slicesEqual(path, exp) {
		t.Errorf("expected path %v got %v", exp, path)
	}
}
