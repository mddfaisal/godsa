package bst

import (
	"ds/stack"
	"fmt"
)

type node struct {
	key   int
	left  *node
	right *node
}

type Bst struct {
	Tree *node
}

func NewBst() *Bst {
	return new(Bst)
}

func (b *Bst) Insert(i int) {
	var (
		root *node = b.Tree
		temp *node
		e    = &node{key: i}
	)
	for root != nil {
		temp = root
		if i < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	if temp == nil {
		b.Tree = e
	} else if i < temp.key {
		temp.left = e
	} else {
		temp.right = e
	}

}

func (b *Bst) Delete(i int) {

}

func (b *Bst) Search() bool {
	found := false
	return found
}

func (b *Bst) Max() int {
	return 0
}

func (b *Bst) Min() int {
	return 0
}

func (b *Bst) TreeSuccessor(i int) {

}

func (b *Bst) InorderDisplay() {
	var (
		current = b.Tree
		done    = false
		st      = stack.NewStack(1000)
	)
	fmt.Println("Inorder display")
	for !done {
		if current != nil {
			st.Push(current)
			current = current.left
		} else {
			if !st.IsEmpty() {
				current = st.Pop().(*node)
				fmt.Printf("%d ", current.key)
				current = current.right
			} else {
				done = true
			}
		}
	}
	fmt.Println("")
}
