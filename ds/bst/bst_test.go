package bst_test

import (
	"ds/bst"
	"testing"
)

var (
	arr         = []int{1, 6, 12, 9, 56, 65, 23, 37, 44, 19, 45, 35, 29, 61, 39}
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
	b.Display()
}
