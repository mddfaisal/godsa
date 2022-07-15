package jumpsearch_test

import (
	"algorithms/searching/jumpsearch"
	"algorithms/sorting/mergesort"
	"algorithms/utils"
	"testing"
	"time"
)

func TestJumpSearch(t *testing.T) {
	t.Log("Jump search...")
	var (
		i     = 807151263
		arr   = mergesort.Sort(utils.LongInputArr(jumpsearch.Path))
		now   = time.Now()
		found = jumpsearch.JumpSearch(arr, i)
	)
	t.Log("Key found: ", found)
	t.Log("Jump search time: ", time.Since(now))
}
