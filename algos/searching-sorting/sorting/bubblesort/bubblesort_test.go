package bubblesort_test

import (
	"searching-sorting/sorting/bubblesort"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	t.Log("Bubble Sort...")
	now := time.Now()
	bubblesort.Sort(utils.LongInputArr(bubblesort.Path))
	t.Log("Elapsed...", time.Since(now))
}
