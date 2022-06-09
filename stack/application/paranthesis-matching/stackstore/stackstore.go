package stackstore

import "fmt"

const (
	OPEN_FLOWER_BRACKET   = "{"
	CLOSED_FLOWER_BRACKET = "}"
	OPEN_ROUND_BRACKET    = "("
	CLOSED_ROUND_BRACKET  = ")"
	OPEN_SQUARE_BRACKET   = "["
	CLOSED_SQUARE_BRACKET = "]"
)

func Mapping(i string) string {
	switch i {
	case OPEN_FLOWER_BRACKET:
		return CLOSED_FLOWER_BRACKET
	case OPEN_SQUARE_BRACKET:
		return CLOSED_SQUARE_BRACKET
	case OPEN_ROUND_BRACKET:
		return CLOSED_ROUND_BRACKET
	default:
		return ""
	}
}

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

func (s *Stack) Push(i string) {
	if s.IsFull() {
		return
	}
	if s.IsEmpty() {
		s.Store = &Node{Data: i}
		s.Top++
		return
	}
	e := &Node{Data: i, Next: nil}
	e.Next = s.Store
	s.Store = e
	s.Top++
}

func (s *Stack) Pop() string {
	r := ""
	if s.IsEmpty() {
		return r
	}
	r = s.Store.Data
	s.Store = s.Store.Next
	s.Top--
	return r
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return s.Store.Data
}

func (s *Stack) IsFull() bool {
	return s.Size == s.Top+1
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Display() {
	fmt.Println("")
	store := s.Store
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("%s -> ", store.Data)
		store = store.Next
	}
	fmt.Println("")
}
