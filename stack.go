package main

import (
	"fmt"
)

type Stack struct {
	data []interface{}
	size int
}

func New() *Stack {
	s := &Stack{
		data: make([]interface{}, 0, 0),
		size: 0,
	}
	return s
}

func (s *Stack) Push(e interface{}) {
	s.data = append(s.data, e)
	s.size++
}

func (s *Stack) Pop() interface{} {
	s.size--
	r := s.data[s.size]
	s.data = s.data[:s.size]

	return r
}

func main() {
	s := New()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	fmt.Println(s.data)
	s.Pop()
	fmt.Println(s.data)
}
