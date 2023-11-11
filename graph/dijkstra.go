package graph

import (
	"container/heap"

	"golang.org/x/exp/slices"
)

type nodeDist struct {
	node int
	dist float64
}

type DistHeap []nodeDist

func (h DistHeap) Len() int           { return len(h) }
func (h DistHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h DistHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DistHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(nodeDist))
}

func (h *DistHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func DijkstraList(graph AdjacencyList, source int, sink int) []int {
	seen := make([]bool, len(graph))
	dists := make([]float64, len(graph))
	h := &DistHeap{}
	heap.Init(h)
	heap.Push(h, nodeDist{node: source, dist: 0})
	for i := range dists {
		dists[i] = infinity
	}
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	dists[source] = 0

	var nd nodeDist
	var dist float64
	var adjs []GraphEdge
	var edge GraphEdge
	for h.Len() > 0 {
		nd = heap.Pop(h).(nodeDist)
		seen[nd.node] = true

		adjs = graph[nd.node]
		for i := 0; i < len(adjs); i++ {
			edge = adjs[i]
			if seen[edge.to] {
				continue
			}

			dist = nd.dist + edge.weight
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				heap.Push(h, nodeDist{node: edge.to, dist: dist})
				prev[edge.to] = nd.node
			}
		}
	}
	out := make([]int, 0)
	curr := sink
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = append(out, source)
	slices.Reverse(out)
	return out
}
