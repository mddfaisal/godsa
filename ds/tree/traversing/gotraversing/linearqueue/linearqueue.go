package linearqueue

import (
	"fmt"
	"gotraversing/structs"
)

type Queue struct {
	Size  int
	Top   int
	Store *structs.QNode
}

func NewQueue(size int) *Queue {
	q := new(Queue)
	q.Size = size
	q.Top = -1
	return q
}

func (q *Queue) Enqueue(i *structs.TNode) {
	if q.IsFull() {
		return
	}
	e := &structs.QNode{Data: i}
	if q.IsEmpty() {
		q.Store = e
		q.Top++
		return
	}
	store := q.Store
	for i := q.Top; i >= 1; i-- {
		store = store.Next
	}
	store.Next = e
	q.Top++
}

func (q *Queue) Dequeue() *structs.TNode {
	r := &structs.TNode{}
	if q.IsEmpty() {
		return r
	}
	r = q.Store.Data
	q.Store = q.Store.Next
	q.Top--
	return r
}

func (q *Queue) IsEmpty() bool {
	return q.Top == -1
}

func (q *Queue) IsFull() bool {
	return q.Size == q.Top+1
}

func (q *Queue) Display() {
	store := q.Store
	fmt.Println("")
	for i := q.Top; i >= 0; i-- {
		fmt.Printf("(%p, (%p, %s, %p)) -> ", store.Data, store.Data.Lchild, store.Data.Data, store.Data.Rchild)
		store = store.Next
	}
	fmt.Println("")
}
