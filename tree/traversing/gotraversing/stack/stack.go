package stack

import (
	"fmt"
	"gotraversing/structs"
)

type Stack struct {
	Top   int
	Size  int
	Store *structs.QNode
}

func NewStack(size int) *Stack {
	s := new(Stack)
	s.Top = -1
	s.Size = size
	return s
}

func (s *Stack) Push(n *structs.TNode) {
	if s.IsFull() {
		return
	}
	if s.IsEmpty() {
		s.Store = &structs.QNode{Data: n, Next: nil}
		s.Top++
		return
	}
	e := &structs.QNode{Data: n, Next: nil}
	e.Next = s.Store
	s.Store = e
	s.Top++
}

func (s *Stack) Pop() *structs.TNode {
	n := new(structs.TNode)
	if s.IsEmpty() {
		return n
	}
	n = s.Store.Data
	s.Store = s.Store.Next
	s.Top--
	return n
}

func (s *Stack) Peek() *structs.TNode {
	if s.IsEmpty() {
		return new(structs.TNode)
	}
	return s.Store.Data
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) IsFull() bool {
	return s.Size == s.Top+1
}

func (s *Stack) Display() {
	store := s.Store
	fmt.Println("")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("(%p, (%p, %s, %p)) -> ", store.Data, store.Data.Lchild, store.Data.Data, store.Data.Rchild)
		store = store.Next
	}
	fmt.Println("")
}
