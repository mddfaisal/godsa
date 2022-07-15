package stack

import (
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type Stack struct {
	Top   int
	Size  int
	Store *node
}

func NewStack(size int) *Stack {
	s := new(Stack)
	s.Size = size
	s.Top = -1
	return s
}

func (s *Stack) Push(i interface{}) {
	if s.IsFull() {
		return
	}
	if s.IsEmpty() {
		s.Store = &node{data: i, next: nil}
		s.Top++
		return
	}
	e := &node{data: i, next: nil}
	e.next = s.Store
	s.Store = e
	s.Top++
}

func (s *Stack) Pop() interface{} {
	var r interface{} = 1
	if s.IsEmpty() {
		return r
	}
	r = s.Store.data
	s.Store = s.Store.next
	s.Top--
	return r
}

func (s *Stack) Peek() interface{} {
	return s.Store.data
}

func (s *Stack) IsFull() bool {
	return s.Top+1 == s.Size
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Display() {
	store := s.Store
	fmt.Println("")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("%v -> ", store.data)
		store = store.next
	}
	fmt.Println("")
}
