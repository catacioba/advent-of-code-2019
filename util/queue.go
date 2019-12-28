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
