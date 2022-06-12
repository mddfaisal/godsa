package stackstore

type Node struct {
	Data string
	Next *Node
}

type Stack struct {
	Top   int
	Size  int
	Store *Node
}

func NewStack(size int) *Stack {
	s := new(Stack)
	s.Size = size
	s.Top = -1
	return s
}

func (s *Stack) IsFull() bool {
	return s.Top+1 == s.Size
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Push(i string) {
	if s.IsFull() {
		return
	}
	e := &Node{Data: i}
	s.Top++
	if s.IsEmpty() {
		s.Store = e
		return
	}
	e.Next = s.Store
	s.Store = e
}

func (s *Stack) Pop() string {
	r := ""
	if s.IsEmpty() {
		return r
	}
	s.Top--
	r = s.Store.Data
	s.Store = s.Store.Next
	return r
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return s.Store.Data
}
