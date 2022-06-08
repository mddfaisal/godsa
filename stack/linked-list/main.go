package main

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	top   int
	size  int
	store *node
}

func NewStack(size int) *stack {
	s := new(stack)
	s.top = -1
	s.size = size
	return s
}

func (s *stack) Push(i int) {
	if s.isFull() {
		return
	}
	if s.isEmpty() {
		s.store = &node{data: i, next: nil}
		s.top++
		return
	}
	e := &node{data: i, next: nil}
	e.next = s.store
	s.store = e
	s.top++
}

func (s *stack) Pop() int {
	r := -1
	if s.isEmpty() {
		return r
	}
	r = s.store.data
	s.store = s.store.next
	s.top--
	return r
}

func (s *stack) Peek() int {
	if s.isEmpty() {
		return -1
	}
	return s.store.data
}

func (s *stack) isFull() bool {
	return s.size == s.top+1
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) Display() {
	store := s.store
	fmt.Println("")
	for i := s.top; i >= 0; i-- {
		fmt.Printf("%d -> ", store.data)
		store = store.next
	}
	fmt.Println("")
}

func main() {
	a := []int{9, 3, 2, 7, 5, 4}
	st := NewStack(6)
	for _, val := range a {
		st.Push(val)
	}
	st.Display()
	fmt.Println("peek", st.Peek())
	fmt.Println(st.Pop())
	st.Display()
	fmt.Println("peek", st.Peek())
}
