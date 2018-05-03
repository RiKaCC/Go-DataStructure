package main

import (
	"fmt"
)

func NewStack() *Stack {
	s := &Stack{
		data: make([]interface{}, 0),
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

func (s *Stack) Empty() bool {
	if s.size == 0 {
		return true
	}

	return false
}

func (s *Stack) SetEmpty() {
	tempArr := make([]interface{}, 0)
	s.data = tempArr
	s.size = 0
}
