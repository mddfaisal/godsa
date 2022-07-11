package exponentialsearch_test

import (
	"searching-sorting/searching/exponentialsearch"
	"searching-sorting/searching/jumpsearch"
	"searching-sorting/sorting/mergesort"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestExponentialSearch(t *testing.T) {
	t.Log("Exponential search...")
	var (
		i     = 807151263
		arr   = mergesort.Sort(utils.LongInputArr(jumpsearch.Path))
		now   = time.Now()
		found = exponentialsearch.Search(arr, i)
	)
	t.Log("Key found: ", found)
	t.Log("Exponential search time: ", time.Since(now))
}
