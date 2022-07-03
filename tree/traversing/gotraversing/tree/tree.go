package tree

import (
	"fmt"
	"gotraversing/linearqueue"
	"gotraversing/structs"
)

type Tree struct {
	Store *structs.TNode
	q     *linearqueue.Queue
	count int
}

func NewTree(q *linearqueue.Queue) *Tree {
	t := new(Tree)
	t.q = q
	return t
}

func (t *Tree) CreateTree() {
	if t.count != 0 {
		fmt.Println("unable to create tree.")
		return
	}
	var x int
	fmt.Printf("Enter root value: ")
	fmt.Scanf("%d", &x)
	t.Store = &structs.TNode{Data: x}
	t.q.Enqueue(t.Store)

	for !t.q.IsEmpty() {
		p := t.q.Dequeue()
		fmt.Printf("\nEnter left child of %d: ", p.Data)
		fmt.Scanf("%d", &x)
		if x != -1 {
			temp := new(structs.TNode)
			temp.Data = x
			temp.Lchild = nil
			temp.Rchild = nil
			p.Lchild = temp
			t.q.Enqueue(temp)
		}

		fmt.Printf("\nEnter right child of %d: ", p.Data)
		fmt.Scanf("%d", &x)
		if x != -1 {
			temp := new(structs.TNode)
			temp.Data = x
			temp.Lchild = nil
			temp.Rchild = nil
			p.Rchild = temp
			t.q.Enqueue(temp)
		}
	}
}
