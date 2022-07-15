package bubblesort_test

import (
	"godsa/algorithms/sorting/bubblesort"
	"godsa/utils"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	t.Log("Bubble Sort...")
	now := time.Now()
	bubblesort.Sort(utils.LongInputArr(bubblesort.Path))
	t.Log("Elapsed...", time.Since(now))
}
