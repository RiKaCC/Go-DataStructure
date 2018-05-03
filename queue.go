package main

type Queue struct {
	data []interface{}
	size int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
		size: 0,
	}
}

func (q *Queue) Push(e interface{}) {
	q.data = append(q.data, e)
	q.size++
}

func (q *Queue) Pop() interface{} {
	r := q.data[0]
	q.data = q.data[1:q.size]
	q.size--
	return r
}

func (q *Queue) Empty() bool {
	if q.size == 0 {
		return true
	}

	return false
}

func (q *Queue) SetEmpty() {
	tempArr := make([]interface{}, 0)
	q.data = tempArr
	q.size = 0
}
