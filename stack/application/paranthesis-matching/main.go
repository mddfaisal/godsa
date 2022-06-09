package main

import (
	"fmt"
	"paranthesis-matching/stackstore"
)

func main() {
	exp := "{[a+(c+d)]+(e+[f+g])}]"
	expBytes := []byte(exp)
	st := stackstore.NewStack(100)
	for _, v := range expBytes {
		s := string(v)
		if s == stackstore.OPEN_FLOWER_BRACKET || s == stackstore.OPEN_ROUND_BRACKET || s == stackstore.OPEN_SQUARE_BRACKET {
			st.Push(s)
		} else if s == stackstore.CLOSED_FLOWER_BRACKET || s == stackstore.CLOSED_ROUND_BRACKET || s == stackstore.CLOSED_SQUARE_BRACKET {
			if st.IsEmpty() {
				st.Push(s)
				continue
			}
			if s == stackstore.Mapping(st.Peek()) {
				st.Pop()
			}
		}
	}
	st.Display()
	if st.IsEmpty() {
		fmt.Println("expression: " + exp + " is balanced")
	} else {
		fmt.Println("expression: " + exp + " is not balanced")
	}
}
