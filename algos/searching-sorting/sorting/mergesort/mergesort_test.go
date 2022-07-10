package mergesort_test

import (
	"searching-sorting/sorting/mergesort"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	t.Log("Merge Sort...")
	var (
		arr = utils.LongInputArr(mergesort.Path)
		now = time.Now()
	)
	mergesort.Sort(arr)
	t.Log("Merge sort time: ", time.Since(now))
}
