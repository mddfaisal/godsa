package structs

type TNode struct {
	Lchild *TNode
	Rchild *TNode
	Data   string
}

type QNode struct {
	Data *TNode
	Next *QNode
}
