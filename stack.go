package main

type Stack struct {
	data []interface{}
	size int
}

func NewStack() *Stack {
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
