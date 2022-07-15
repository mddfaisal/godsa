package stack

import (
	"fmt"
	"godsa/datatype"
	"godsa/utils"
)

type Stack datatype.Stack

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
		s.Store = &datatype.Node{Data: i, Next: nil}
		s.Top++
		return
	}
	e := &datatype.Node{Data: i, Next: nil}
	e.Next = s.Store
	s.Store = e
	s.Top++
}

func (s *Stack) Pop() interface{} {
	var r interface{} = 1
	if s.IsEmpty() {
		return r
	}
	r = s.Store.Data
	s.Store = s.Store.Next
	s.Top--
	return r
}

func (s *Stack) Peek() interface{} {
	return s.Store.Data
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
		fmt.Printf("%v -> ", utils.Interface2Val(store.Data))
		store = store.Next
	}
	fmt.Println("")
}
