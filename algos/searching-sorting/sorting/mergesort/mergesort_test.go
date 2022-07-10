package mergesort_test

import (
	"searching-sorting/sorting/mergesort"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	t.Log("Merge Sort...")
	now := time.Now()
	mergesort.Sort(utils.LongInputArr(mergesort.Path))
	t.Log("Merge sort time: ", time.Since(now))
}
