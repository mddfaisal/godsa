package main

import "fmt"

type stack struct {
	top  int
	size int
	arr  []int
}

func NewStack(size int) *stack {
	s := new(stack)
	s.top = -1
	s.size = size
	s.arr = []int{}
	return s
}

func (s *stack) Display() {
	fmt.Println("")
	for i := s.top; i >= 0; i-- {
		fmt.Printf("%d ", s.arr[i])
	}
	fmt.Println("")
}

func (s *stack) Push(i int) {
	if s.isFull() {
		return
	}
	s.top++
	s.arr = append(s.arr, i)
}

func (s *stack) Pop() int {
	r := -1
	if s.isEmpty() {
		return r
	}
	r = s.arr[s.top]
	s.arr = s.arr[:s.top]
	s.top--
	return r
}

func (s *stack) Peek() int {
	r := -1
	if s.isEmpty() {
		return r
	}
	return s.arr[s.top]
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) isFull() bool {
	return s.size == s.top+1
}

func main() {
	a := []int{9, 3, 2, 7, 5, 4}
	st := NewStack(6)
	for _, val := range a {
		st.Push(val)
	}
	fmt.Println(st.Peek())
	st.Display()
	st.Pop()
	st.Display()
	fmt.Println(st.Peek())
}
