package main

import (
	"fmt"
)

type Queue struct {
	data []interface{}
	size int
}

func New() *Queue {
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

func main() {
	q := New()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	fmt.Println(q.data)
	fmt.Println(q.Pop())
	fmt.Println(q.data)
}
