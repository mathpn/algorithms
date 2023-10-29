package main

import "testing"

func createTree() *BinaryTree[int] {
	return &BinaryTree[int]{
		root: &bnode[int]{
			value: 20,
			right: &bnode[int]{
				value: 50,
				right: &bnode[int]{
					value: 100,
					right: nil,
					left:  nil,
				},
				left: &bnode[int]{
					value: 30,
					right: &bnode[int]{
						value: 45,
						right: nil,
						left:  nil,
					},
					left: &bnode[int]{
						value: 29,
						right: nil,
						left:  nil,
					},
				},
			},
			left: &bnode[int]{
				value: 10,
				right: &bnode[int]{
					value: 15,
					right: nil,
					left:  nil,
				},
				left: &bnode[int]{
					value: 5,
					right: &bnode[int]{
						value: 7,
						right: nil,
						left:  nil,
					},
					left: nil,
				},
			},
		},
	}
}

func createTree2() *BinaryTree[int] {
	return &BinaryTree[int]{
		root: &bnode[int]{
			value: 20,
			right: &bnode[int]{
				value: 50,
				right: nil,
				left: &bnode[int]{
					value: 30,
					right: &bnode[int]{
						value: 45,
						right: &bnode[int]{
							value: 49,
							right: nil,
							left:  nil,
						},
						left: nil,
					},
					left: &bnode[int]{
						value: 29,
						right: nil,
						left: &bnode[int]{
							value: 21,
							right: nil,
							left:  nil,
						},
					},
				},
			},
			left: &bnode[int]{
				value: 10,
				right: &bnode[int]{
					value: 15,
					right: nil,
					left:  nil,
				},
				left: &bnode[int]{
					value: 5,
					right: &bnode[int]{
						value: 7,
						right: nil,
						left:  nil,
					},
					left: nil,
				},
			},
		},
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

func TestBinaryTree(t *testing.T) {
	t.Run("pre-order traversal", func(t *testing.T) {
		tree := createTree()
		path := tree.PreOrderTraversal()
		exp := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}

		if !slicesEqual(exp, path) {
			t.Errorf("expected %v, got %v", exp, path)
		}
	})
	t.Run("post-order traversal", func(t *testing.T) {
		tree := createTree()
		path := tree.PostOrderTraversal()
		exp := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}

		if !slicesEqual(exp, path) {
			t.Errorf("expected %v, got %v", exp, path)
		}
	})
	t.Run("in-order traversal", func(t *testing.T) {
		tree := createTree()
		path := tree.InOrderTraversal()
		exp := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}

		if !slicesEqual(exp, path) {
			t.Errorf("expected %v, got %v", exp, path)
		}
	})
	t.Run("breadth-first traversal", func(t *testing.T) {
		tree := createTree()
		path := tree.BreadthFirstTraversal()
		exp := []int{20, 10, 50, 5, 15, 30, 100, 7, 29, 45}

		if !slicesEqual(exp, path) {
			t.Errorf("expected %v, got %v", exp, path)
		}
	})
	t.Run("breadth-first search", func(t *testing.T) {
		tree := createTree()
		values := []int{20, 15, 30}
		for _, v := range values {
			found := tree.BreadthFirstSearch(v)
			if !found {
				t.Errorf("value %d not found in tree", v)
			}
		}
		values = []int{-1, 42, 101010}
		for _, v := range values {
			found := tree.BreadthFirstSearch(v)
			if found {
				t.Errorf("value %d should not be found in tree", v)
			}
		}
	})
}
