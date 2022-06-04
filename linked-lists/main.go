package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	List   *Node
	Length int
}

func NewNode() *LinkedList {
	return new(LinkedList)
}

func (n *LinkedList) AddElement(data int) {
	if n.Length == 0 {
		n1 := new(Node)
		n1.Data = data
		n.List = n1
	} else {
		n1 := new(Node)
		n1.Data = data
		n1.Next = n.List
		n.List = n1
	}
	n.Length++
}

func (n *LinkedList) Display() {
	list := n.List
	i := 0
	fmt.Println("")
	for {
		fmt.Printf("(%d, %d) -> ", list.Data, i)
		if list.Next == nil {
			break
		}
		list = list.Next
		i++
	}
	fmt.Println("")
}

func (n *LinkedList) Search(data int) {
	list := n.List
	for {
		if list.Data == data {
			fmt.Println("found: ", data)
			return
		}
		if list.Next == nil {
			break
		}
		list = list.Next
	}
	fmt.Println("not found: ", data)
}

func (n *LinkedList) Insert(data int, index int) {
	list := n.List
	e := new(Node)
	e.Data = data
	if index == 0 {
		e.Next = n.List
		n.List = e
	} else {
		for i := 0; i < index-1; i++ {
			list = list.Next
		}
		e.Next = list.Next
		list.Next = e
	}
	n.Length++
}

func (n *LinkedList) Remove(index int) {
	list := n.List
	for i := 0; i < index-1; i++ {
		list = list.Next
	}
	list.Next = list.Next.Next
	n.Length--
}

func (n *LinkedList) Sort() {
	fmt.Println("length: ", n.Length)
	// list := n.List
	// for i := 0; i <= n.Length; i++ {
	// 	for j := 0; j < n.Length-i-1; j++ {

	// 	}
	// }
}

func main() {
	a := []int{3, 6, 7, 8, 1}
	n := NewNode()
	for i := len(a) - 1; i >= 0; i-- {
		n.AddElement(a[i])
	}
	n.Search(99)
	n.Search(67)
	n.Display()

	n.Insert(99, 3)
	n.Display()

	n.Remove(3)
	n.Display()

	n.Sort()
	n.Display()
}