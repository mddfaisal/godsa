package jumpsearch_test

import (
	"searching-sorting/searching/jumpsearch"
	"searching-sorting/sorting/mergesort"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestJumpSearch(t *testing.T) {
	t.Log("Jump search...")
	var (
		i     = 807151263
		arr   = mergesort.Sort(utils.LongInputArr(jumpsearch.Path), i)
		now   = time.Now()
		found = jumpsearch.JumpSearch(arr)
	)
	t.Log("Key found: ", found)
	t.Log("Jump search time: ", time.Since(now))
}
