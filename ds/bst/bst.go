package bst

import (
	"ds/datatype"
	"ds/stack"
	"fmt"
)

type Bst datatype.Bst

func NewBst() *Bst {
	return new(Bst)
}

func (b *Bst) Insert(i int) {
	var (
		root = b.Tree
		temp *datatype.TNode
	)
	for root != nil {
		temp = root
		if i < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	if temp == nil {
		b.Tree = &datatype.TNode{Key: i}
	} else if i < temp.Key {
		temp.Left = &datatype.TNode{Key: i, Parent: temp}
	} else {
		temp.Right = &datatype.TNode{Key: i, Parent: temp}
	}

}

func (b *Bst) Delete(i int) {

}

func (b *Bst) Transplant(u *datatype.TNode, v *datatype.TNode) {
	if u.Parent == nil {
		b.Tree = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

func (b *Bst) Search(i int) bool {
	var (
		root = b.Tree
	)
	for root != nil && root.Key != i {
		if i > root.Key {
			root = root.Right
			continue
		}
		root = root.Left
	}
	if root == nil {
		return false
	}
	return root.Key == i
}

func (b *Bst) Max() int {
	root := b.Tree
	for root.Right != nil {
		root = root.Right
	}
	return root.Key
}

func (b *Bst) Min() int {
	root := b.Tree
	for root.Left != nil {
		root = root.Left
	}
	return root.Key
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
			current = current.Left
		} else {
			if !st.IsEmpty() {
				current = st.Pop().(*datatype.TNode)
				fmt.Printf("%d -> ", current.Key)
				current = current.Right
			} else {
				done = true
			}
		}
	}
	fmt.Println("")
}

func (b *Bst) PreorderDisplay() {
	var (
		stk  = stack.NewStack(100)
		root = b.Tree
	)
	fmt.Println("Preorder display")
	for !stk.IsEmpty() || root != nil {
		if root == nil {
			root = stk.Pop().(*datatype.TNode)
			root = root.Right
			continue
		}
		fmt.Printf("%d -> ", root.Key)
		stk.Push(root)
		root = root.Left
	}
	fmt.Println("")
}

func (b *Bst) PostorderDisplay() {
	var (
		xstk = stack.NewStack(100)
		ystk = stack.NewStack(100)
		root = b.Tree
	)
	xstk.Push(root)
	fmt.Println("Postorder display")
	for !xstk.IsEmpty() {
		n := xstk.Pop().(*datatype.TNode)
		ystk.Push(n)
		if n.Left != nil {
			xstk.Push(n.Left)
		}
		if n.Right != nil {
			xstk.Push(n.Right)
		}
	}
	ystk.Display()
	fmt.Println("")
}

func (b *Bst) Display() {
	fmt.Println(b.Tree.Key, b.Tree.Parent)
}
