package stack_test

import (
	"fmt"
	"godsa/ds/stack"
	"testing"
)

func TestStack(t *testing.T) {
	a := []int{9, 3, 2, 7, 5, 4}
	st := stack.NewStack(10)
	for _, val := range a {
		st.Push(val)
	}
	st.Display()
	fmt.Println("peek", st.Peek())
	fmt.Println(st.Pop())
	st.Display()
	fmt.Println("peek", st.Peek())
}
