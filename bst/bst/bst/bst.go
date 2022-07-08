package bst

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Bst struct {
	Store *Node
}

func NewBst() *Bst {
	return new(Bst)
}

func (b *Bst) Insert() {

}

func (b *Bst) Search() {

}

func (b *Bst) Delete(i int) {

}

func (b *Bst) PreOrderDisplay() {

}
