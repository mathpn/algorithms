package tree

import (
	"sort"
	"testing"
)

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
	t.Run("depth-first search", func(t *testing.T) {
		tree := createTree()
		values := []int{20, 15, 30}
		for _, v := range values {
			found := tree.DepthFirstSearch(v)
			if !found {
				t.Errorf("value %d not found in tree", v)
			}
		}
		values = []int{-1, 42, 101010}
		for _, v := range values {
			found := tree.DepthFirstSearch(v)
			if found {
				t.Errorf("value %d should not be found in tree", v)
			}
		}
	})
	t.Run("tree comparison", func(t *testing.T) {
		a := createTree()
		b := createTree2()
		c := createTree()
		if CompareTrees(a, b) {
			t.Errorf("different trees considered equal: %v - %v", a.InOrderTraversal(), b.InOrderTraversal())
		}
		if !CompareTrees(a, a) {
			t.Errorf("tree should be equal to itself: %v", a.InOrderTraversal())
		}
		if !CompareTrees(b, b) {
			t.Errorf("tree should be equal to itself: %v", b.InOrderTraversal())
		}
		if !CompareTrees(a, c) {
			t.Errorf("tree should be equal to identical tree: %v - %v", a.InOrderTraversal(), c.InOrderTraversal())
		}
	})
}

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()
	values := []int{5, 2, 70, 42, 4, 1, 8, 7}

	for _, v := range values {
		heap.Insert(v)
	}
	sort.Ints(values)
	for i, v := range values {
		hv, err := heap.Delete()
		if err != nil {
			t.Fatal(err)
		}
		if hv != v {
			t.Errorf("expected %d, got %d", v, hv)
		}
		expLen := len(values) - i - 1
		if heap.Len != expLen {
			t.Errorf("expected length of %d, got %d", expLen, heap.Len)
		}
	}
}

func TestBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree()

	var err error
	_, err = tree.Search(-1)
	if err == nil {
		t.Errorf("expected error when searching in empty tree")
	}

	keys := []int{5, 2, 70, 42, 4, 1, 8, 7}
	values := []string{"foo", "bar", "ping", "pong", "hey", "search", "tree", "binary"}

	for i := 0; i < len(keys); i++ {
		tree.Insert(keys[i], values[i])
	}

	indices := make([]int, len(keys))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return keys[indices[i]] < keys[indices[j]]
	})

	var out string
	for i := 0; i < len(keys); i++ {
		idx := indices[i]
		out, err = tree.Search(keys[idx])
		if err != nil {
			t.Fatal(err)
		}
		exp := values[idx]
		if out != exp {
			t.Errorf("expected '%s', got '%s'", exp, out)
		}
	}

	_, err = tree.Search(-1)
	if err == nil {
		t.Error("expected error when searching for non-existing key")
	}
}
