package tree

import (
	"fmt"
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

type treeInsert struct {
	value string
	key   int
}

func makeInserts() [][]treeInsert {
	out := make([][]treeInsert, 3)
	out = append(out,
		[]treeInsert{
			{key: 5, value: "foo"},
			{key: 2, value: "bar"},
			{key: 70, value: "ping"},
			{key: 42, value: "pong"},
			{key: 4, value: "hey"},
			{key: 1, value: "search"},
			{key: 8, value: "tree"},
			{key: 7, value: "binary"},
		})

	inserts := make([]treeInsert, 0)
	for i := 0; i < 10; i++ {
		inserts = append(inserts, treeInsert{key: i, value: fmt.Sprintf("v%d", i)})
	}
	out = append(out, inserts)

	inserts = make([]treeInsert, 0)
	for i := 0; i < 10; i++ {
		inserts = append(inserts, treeInsert{key: 10 - i, value: fmt.Sprintf("v%d", 10-i)})
	}
	out = append(out, inserts)
	return out
}

type treeConstructor func() BinarySearchable[int, string]

func testBinarysearchable(t *testing.T, constructor treeConstructor, inserts []treeInsert) {
	tree := constructor()
	var err error
	_, err = tree.Search(-1)
	if err == nil {
		t.Errorf("expected error when searching in empty tree")
	}

	var treeDump string
	for i, insert := range inserts {
		tree.Insert(insert.key, insert.value)
		treeDump, err = tree.Dump()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("step %d -> tree structure:\n%s", i+1, treeDump)
	}

	var out string
	for _, insert := range inserts {
		out, err = tree.Search(insert.key)
		if err != nil {
			t.Fatal(err)
		}
		if out != insert.value {
			t.Errorf("expected '%s', got '%s'", insert.value, out)
		}
	}

	_, err = tree.Search(-1)
	if err == nil {
		t.Error("expected error when searching for non-existing key")
	}
}

func TestBinarySearchTree(t *testing.T) {
	insertsSlice := makeInserts()
	for _, inserts := range insertsSlice {
		testBinarysearchable(t, NewBinarySearchTree[int, string], inserts)
	}
}

func TestAVLTree(t *testing.T) {
	insertsSlice := makeInserts()
	for _, inserts := range insertsSlice {
		testBinarysearchable(t, NewAVLTree[int, string], inserts)
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()
	inserts := []string{"cat", "can", "foo", "the", "then", "breathe"}
	var found bool
	for _, word := range inserts {
		found = trie.Search(word)
		if found {
			t.Errorf("word %s should not be found in trie", word)
		}
		trie.Insert(word)

		found = trie.Search(word)
		if !found {
			t.Errorf("word %s should be found in trie", word)
		}
	}

	for _, word := range inserts {
		found = trie.Search(word)
		if !found {
			t.Errorf("word %s should be found in trie", word)
		}
		trie.Insert(word)
	}
}

type prefixTest struct {
	word   string
	prefix bool
	insert bool
}

func TestTriePrefix(t *testing.T) {
	trie := NewTrie()
	tests := []prefixTest{
		{word: "ca", prefix: false, insert: false},
		{word: "c", prefix: false, insert: false},
		{word: "cat", prefix: false, insert: true},
		{word: "can", prefix: false, insert: true},
		{word: "ca", prefix: true, insert: false},
		{word: "the", prefix: false, insert: true},
		{word: "then", prefix: false, insert: true},
		{word: "the", prefix: true, insert: true},
		{word: "the", prefix: true, insert: true},
	}

	var isPrefix bool
	for _, prefixTest := range tests {
		isPrefix = trie.StartsWith(prefixTest.word)
		if isPrefix != prefixTest.prefix {
			t.Errorf(
				"trie prefix search failed for word %s. Expected %v got %v",
				prefixTest.word,
				prefixTest.prefix,
				isPrefix,
			)
		}
		if prefixTest.insert {
			trie.Insert(prefixTest.word)
		}
	}
}

func TestPatriciaTrie(t *testing.T) {
	trie := NewPatriciaTrie()
	inserts := []string{
		"orange",
		"organism",
		"apple",
		"ape",
		"cat",
		"can",
		"foo",
		"the",
		"then",
		"bar",
		"organization",
		"organizations",
		"oranges",
		"organized",
		"organs",
		"horror",
		"ore",
		"oregon",
	}
	var found bool
	for _, word := range inserts {
		found = trie.Search(word)
		if found {
			t.Errorf("word %s should not be found in trie", word)
		}
		trie.Insert(word)
		trie.Print()

		found = trie.Search(word)
		if !found {
			t.Errorf("word %s should be found in trie", word)
		}
	}

	for _, word := range inserts {
		found = trie.Search(word)
		if !found {
			t.Errorf("word %s should be found in trie", word)
		}
		trie.Insert(word)
	}
}
