package tree

import (
	"fmt"
	"gotraversing/linearqueue"
	"gotraversing/stack"
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
	n := &structs.TNode{Data: "A", Lchild: nil, Rchild: nil}
	t.Store = n
	t.Store.Lchild = &structs.TNode{Data: "B", Lchild: nil, Rchild: nil}
	t.Store.Rchild = &structs.TNode{Data: "C", Lchild: nil, Rchild: nil}

	t.Store.Lchild.Lchild = &structs.TNode{Data: "D", Lchild: nil, Rchild: nil}
	t.Store.Lchild.Rchild = &structs.TNode{Data: "E", Lchild: nil, Rchild: nil}

	t.Store.Rchild.Lchild = &structs.TNode{Data: "F", Lchild: nil, Rchild: nil}
	t.Store.Rchild.Rchild = &structs.TNode{Data: "G", Lchild: nil, Rchild: nil}
	// if t.count != 0 {
	// 	fmt.Println("unable to create tree.")
	// 	return
	// }
	// var x string
	// fmt.Printf("Enter root value: ")
	// fmt.Scanf("%s", &x)
	// t.Store = &structs.TNode{Data: x}
	// t.q.Enqueue(t.Store)

	// for !t.q.IsEmpty() {
	// 	p := t.q.Dequeue()
	// 	fmt.Printf("\nEnter left child of %s: ", p.Data)
	// 	fmt.Scanf("%s", &x)
	// 	if x != "X" {
	// 		temp := new(structs.TNode)
	// 		temp.Data = x
	// 		temp.Lchild = nil
	// 		temp.Rchild = nil
	// 		p.Lchild = temp
	// 		t.q.Enqueue(temp)
	// 	}

	// 	fmt.Printf("\nEnter right child of %s: ", p.Data)
	// 	fmt.Scanf("%s", &x)
	// 	if x != "X" {
	// 		temp := new(structs.TNode)
	// 		temp.Data = x
	// 		temp.Lchild = nil
	// 		temp.Rchild = nil
	// 		p.Rchild = temp
	// 		t.q.Enqueue(temp)
	// 	}
	// }
}

func (t *Tree) Preorder() {
	stk := stack.NewStack(10)
	n := t.Store
	for !stk.IsEmpty() || n != nil {
		if n == nil {
			n = stk.Pop()
			n = n.Rchild
			continue
		}
		fmt.Printf("%s ", n.Data)
		stk.Push(n)
		n = n.Lchild
	}
}

func (t *Tree) Postorder() {}

// 1) Create an empty stack S.
// 2) Initialize current node as root
// 3) Push the current node to S and set current = current->left until current is NULL
// 4) If current is NULL and stack is not empty then
//      a) Pop the top item from stack.
//      b) Print the popped item, set current = popped_item->right
//      c) Go to step 3.
// 5) If current is NULL and stack is empty then we are done.
// https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion/
func (t *Tree) Inorder() {
	var (
		current *structs.TNode = t.Store // set current to root
		st      *stack.Stack   = stack.NewStack(10)
		done    bool           = false
	)
	for !done {
		if current != nil {
			st.Push(current)
			current = current.Lchild
		} else {
			if !st.IsEmpty() {
				current = st.Pop()
				fmt.Printf("%s ", current.Data)
				current = current.Rchild
			} else {
				done = true
			}
		}
	}
}

// 1) Create an empty queue q
// 2) temp_node = root /*start from root*/
// 3) Loop while temp_node is not NULL
//     a) print temp_node->data.
//     b) Enqueue temp_node’s children
//       (first left then right children) to q
//     c) Dequeue a node from q.
// https://www.geeksforgeeks.org/level-order-tree-traversal/?ref=lbp
func (t *Tree) LevelOrder() {
	// step 1
	q := linearqueue.NewQueue(10)

	// step 2
	temp_node := t.Store

	// step 3
	for temp_node.Data != "" {
		fmt.Printf("%s ", temp_node.Data)

		if temp_node.Lchild != nil {
			q.Enqueue(temp_node.Lchild)
		}
		if temp_node.Rchild != nil {
			q.Enqueue(temp_node.Rchild)
		}
		temp_node = q.Dequeue()
	}
}
