package exponentialsearch_test

import (
	"algorithms/searching/exponentialsearch"
	"algorithms/searching/jumpsearch"
	"algorithms/sorting/mergesort"
	"algorithms/utils"
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
