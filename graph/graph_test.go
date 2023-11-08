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
