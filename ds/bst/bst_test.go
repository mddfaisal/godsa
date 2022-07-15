package bst_test

import (
	"ds/bst"
	"testing"
)

func TestInsertionInBst(t *testing.T) {
	t.Log("Creating bst: ")
	b := bst.NewBst()
	for i := 1; i <= 1000; i++ {
		b.Insert(i)
	}
	b.InorderDisplay()
}
