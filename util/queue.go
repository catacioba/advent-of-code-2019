package util

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{
		l: list.New(),
	}
}

func (q *Queue) Push(elem interface{}) {
	q.l.PushBack(elem)
}

func (q *Queue) Pop() interface{} {
	el := q.l.Front()
	val := el.Value
	q.l.Remove(el)
	return val
}

func (q *Queue) Size() int {
	return q.l.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}