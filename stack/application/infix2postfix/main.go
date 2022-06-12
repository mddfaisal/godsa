package main

import (
	"fmt"
	"infix2postfix/expressions"
	"infix2postfix/stackstore"
)

func isbalanced(i string) bool {
	st := stackstore.NewStack(100)
	expBytes := []byte(i)
	for _, v := range expBytes {
		s := string(v)
		if s == expressions.OPEN_FLOWER_BRACKET || s == expressions.OPEN_ROUND_BRACKET || s == expressions.OPEN_SQUARE_BRACKET {
			st.Push(s)
		} else if s == expressions.CLOSED_FLOWER_BRACKET || s == expressions.CLOSED_ROUND_BRACKET || s == expressions.CLOSED_SQUARE_BRACKET {
			if st.IsEmpty() {
				st.Push(s)
				continue
			}
			if s == expressions.Mapping(st.Peek()) {
				st.Pop()
			}
		}
	}
	return st.IsEmpty()
}

func infix2postfix(i string) string {
	return ""
}

func main() {
	fmt.Println("Is balanced?")
	for _, val := range expressions.Exps {
		fmt.Println(expressions.INFIX, ": ", val[expressions.INFIX])
		fmt.Println(expressions.POSTFIX, ": ", val[expressions.POSTFIX])
		fmt.Println("Is balanced: ", isbalanced(val[expressions.INFIX]))
		fmt.Println("")
	}
}
