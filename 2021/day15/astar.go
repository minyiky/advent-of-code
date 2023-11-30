package day15

import "github.com/minyiky/advent-of-code-utils/pkg/point"

type pqItem[P point.Point2D | point.Point3D] struct {
	coords P
	index  int // The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds pqItems.
type PriorityQueue[P point.Point2D | point.Point3D] struct {
	items  []*pqItem[P]
	scores map[P]int
	set    map[P]struct{}
}

func newPriorityQueue[P point.Point2D | point.Point3D](scores map[P]int) *PriorityQueue[P] {
	return &PriorityQueue[P]{
		items:  make([]*pqItem[P], 0),
		scores: scores,
		set:    make(map[P]struct{}),
	}
}

func (pq PriorityQueue[P]) Len() int { return len(pq.items) }

// Decides the priority of items. Here we want to prioritise those with a lower fScore.
// fScore[n] represents our current best guess as to how short a path from start to
// finish can be if it goes through n.
func (pq PriorityQueue[P]) Less(i, j int) bool {
	return pq.scores[pq.items[i].coords] < pq.scores[pq.items[j].coords]
}

func (pq PriorityQueue[P]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

// heap.Push() calls this method
func (pq *PriorityQueue[P]) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*pqItem[P]) // x is the item. My understanding is that the .(*pqItem) on the end is like writing "as pqItem" in TypeScript - the func doesn't know what the interface is otherwise
	item.index = n
	pq.items = append(pq.items, item)
	pq.set[item.coords] = struct{}{}
}

// will be used by heap.Pop() to return the 'minimum' item on the queue based on the Less() func defined above, in our case the one with the lowest value
// therefore our graph won't be traversed by DFS or BFS but simply by which 'neighbours' have the lowest value - this is the efficiency improvement that
// the priority queue offers to our A* algorithm
func (pq *PriorityQueue[P]) Pop() interface{} { // it returns an initialised interface of some sort
	old := *pq
	n := len(old.items)
	item := old.items[n-1]
	old.items[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.items = old.items[0 : n-1]
	delete(pq.set, item.coords)
	return item
}
