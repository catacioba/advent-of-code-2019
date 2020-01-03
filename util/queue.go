package util

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

// Point Queue
type PointQueue struct {
	elems []Point
	idx   int
}

func NewPointQueue() *PointQueue {
	return &PointQueue{
		elems: []Point{},
		idx:   0,
	}
}

func (q *PointQueue) Push(elem Point) {
	q.elems = append(q.elems, elem)
}

func (q *PointQueue) Pop() Point {
	if q.Size() == 0 {
		panic("Pop on empty queue!")
	}
	q.idx++
	return q.elems[q.idx-1]
}

func (q *PointQueue) Size() int {
	return len(q.elems) - q.idx
}
