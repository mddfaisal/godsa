package main

import "fmt"

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

func (n *LinkedList) Insert(data int) *LinkedList {
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
	return n
}

func (n *LinkedList) Display() {
	list := n.List
	for list.Next != nil {
		fmt.Printf("%d -> ", list.Data)
		list = list.Next
		if list.Next == nil {
			fmt.Printf("%d -> ", list.Data)
		}
	}
}

func main() {
	a := []int{3, 6, 7, 8, 1, 4, 9, 5}
	n := NewNode()
	for _, k := range a {
		n.Insert(k)
	}
	n.Display()
}
