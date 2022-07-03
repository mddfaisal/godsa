package structs

type TNode struct {
	Lchild *TNode
	Rchild *TNode
	Data   int
}

type QNode struct {
	Data *TNode
	Next *QNode
}
