package datatype

type Node struct {
	Data interface{}
	Next *Node
}

type Stack struct {
	Top   int
	Size  int
	Store *Node
}

type TNode struct {
	Key    int
	Parent *TNode
	Right  *TNode
	Left   *TNode
}

type Bst struct {
	Tree *TNode
}
