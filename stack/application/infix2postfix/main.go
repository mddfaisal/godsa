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
	postfix := ""
	stk := stackstore.NewStack(100)
	exp := []byte(i)
	for _, val := range exp {
		strVal := string(val)
		// If the scanned character is
		// an operand, add it to output.
		if expressions.IsOperand(strVal) {
			postfix = postfix + strVal
		} else if strVal == expressions.OPEN_ROUND_BRACKET {
			// If the scanned character is an
			// ‘(‘, push it to the stack.
			stk.Push(strVal)
		} else if strVal == expressions.CLOSED_ROUND_BRACKET {
			// If the scanned character is an ‘)’,
			// pop and output from the stack
			// until an ‘(‘ is encountered.
			for !stk.IsEmpty() && stk.Peek() != expressions.OPEN_ROUND_BRACKET {
				postfix = postfix + stk.Pop()
			}
			if !stk.IsEmpty() && stk.Peek() != expressions.OPEN_ROUND_BRACKET {
				return "-1"
			} else {
				stk.Pop()
			}
		} else {
			for !stk.IsEmpty() && expressions.Precedence(strVal) <= expressions.Precedence(stk.Peek()) {
				postfix = postfix + stk.Pop()
			}
			stk.Push(strVal)
		}
	}

	// pop all the operators from the stack
	for !stk.IsEmpty() {
		postfix = postfix + stk.Pop()
	}
	return postfix
}

func main() {
	fmt.Println("Is balanced?")
	for _, val := range expressions.Exps {
		fmt.Println(expressions.INFIX, ": ", val[expressions.INFIX])
		fmt.Println(expressions.POSTFIX, ": ", val[expressions.POSTFIX])
		fmt.Println("Is balanced: ", isbalanced(val[expressions.INFIX]))
		fmt.Println("Postfix result: ", infix2postfix(val[expressions.INFIX]))
		fmt.Println("")
	}
}
