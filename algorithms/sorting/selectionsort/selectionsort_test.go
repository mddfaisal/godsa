package selectionsort_test

import (
	"algorithms/searching/binarysearch"
	"algorithms/sorting/selectionsort"
	"algorithms/utils"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	t.Log("Selection sort...")
	arr := utils.LongInputArr(binarysearch.Path)
	now := time.Now()
	result := selectionsort.Sort(arr)
	t.Log("Selection sort time: ", time.Since(now))
	if !utils.IsSameArr(arr, result) {
		t.Error("Arrays are not same.")
	}
}
