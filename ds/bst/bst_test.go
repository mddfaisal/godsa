package bst_test

import (
	"godsa/ds/bst"
	"testing"
)

var (
	arr         = []int{35, 12, 6, 1, 9, 23, 19, 29, 45, 39, 37, 44, 61, 56, 65}
	search_item = 56
)

func TestInsertionInBst(t *testing.T) {
	t.Log("Creating bst: ")
	b := bst.NewBst()
	for _, i := range arr {
		b.Insert(i)
	}
	t.Log("Max", b.Max())
	t.Log("Min", b.Min())
	t.Log("Searching ", search_item, b.Search(search_item))
	b.InorderDisplay()
	b.PreorderDisplay()
	b.PostorderDisplay()
	b.Display()
}
