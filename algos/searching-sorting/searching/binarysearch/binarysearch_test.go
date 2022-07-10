package binarysearch_test

import (
	"searching-sorting/searching/binarysearch"
	"searching-sorting/sorting/mergesort"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	t.Log("Binary search...")
	arr := mergesort.Sort(utils.LongInputArr(binarysearch.Path))
	i := 807151263
	now := time.Now()
	t.Log("Key found: ", binarysearch.Search(arr, i))
	t.Log("Binary search time: ", time.Since(now))
}
