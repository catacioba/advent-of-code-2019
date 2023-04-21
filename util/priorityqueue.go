package util

import (
	"container/heap"
)

type ArrayPriorityQueue struct {
	inner innerPriorityQueue
}

func NewArrayPriorityQueue() *ArrayPriorityQueue {
	pq := &ArrayPriorityQueue{inner: make([]*pqItem, 0)}
	heap.Init(&pq.inner)
	return pq
}

func (pq *ArrayPriorityQueue) Push(value any, priority int) {
	item := &pqItem{
		value:    value,
		priority: priority,
	}
	heap.Push(&pq.inner, item)
}

func (pq *ArrayPriorityQueue) Pop() any {
	item := heap.Pop(&pq.inner).(*pqItem)
	return item.value
}

func (pq *ArrayPriorityQueue) Empty() bool {
	return pq.inner.Len() == 0
}

type pqItem struct {
	value    any // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap.
}

type innerPriorityQueue []*pqItem

func (pq innerPriorityQueue) Len() int { return len(pq) }

func (pq innerPriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq innerPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *innerPriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *innerPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *innerPriorityQueue) update(item *pqItem, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
