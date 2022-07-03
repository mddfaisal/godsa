package main

import (
	"fmt"
	"gotraversing/linearqueue"
	"gotraversing/structs"
	"gotraversing/tree"
)

func preorder(t *structs.TNode) {
	if t != nil {
		fmt.Printf("%d ", t.Data)
		preorder(t.Lchild)
		preorder(t.Rchild)
	}
}

func postorder(t *structs.TNode) {
	if t != nil {
		postorder(t.Lchild)
		postorder(t.Rchild)
		fmt.Printf("%d ", t.Data)
	}
}

func inorder(t *structs.TNode) {
	if t != nil {
		postorder(t.Lchild)
		fmt.Printf("%d ", t.Data)
		postorder(t.Rchild)
	}
}

func main() {
	t := tree.NewTree(linearqueue.NewQueue(10))
	t.CreateTree()
	fmt.Println("\nPre order: ")
	preorder(t.Store)

	fmt.Println("\nPost order: ")
	postorder(t.Store)

	fmt.Println("\nIn order: ")
	inorder(t.Store)
}
