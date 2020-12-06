package util

import "container/list"

type Queue struct {
	elems []string
	idx   int
}

func NewQueue() *Queue {
	return &Queue{
		elems: []string{},
		idx:   0,
	}
}

func (q *Queue) Push(elem string) {
	q.elems = append(q.elems, elem)
}

func (q *Queue) Pop() string {
	if q.Size() == 0 {
		panic("Pop on empty queue!")
	}
	q.idx++
	return q.elems[q.idx-1]
}

func (q *Queue) Size() int {
	return len(q.elems) - q.idx
}

type PointQueue struct {
	l *list.List
}

func NewPointQueue() *PointQueue {
	return &PointQueue{
		l: list.New(),
	}
}

func (q *PointQueue) Push(elem Point) {
	q.l.PushBack(elem)
}

func (q *PointQueue) Pop() Point {
	el := q.l.Front()
	val := el.Value.(Point)
	q.l.Remove(el)
	return val
}

func (q *PointQueue) Size() int {
	return q.l.Len()
}
