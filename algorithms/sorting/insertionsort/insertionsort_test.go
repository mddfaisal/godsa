package insertionsort_test

import (
	"godsa/algorithms/sorting/insertionsort"
	"godsa/utils"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	t.Log("Insertion Sort")
	now := time.Now()
	insertionsort.Sort(utils.LongInputArr(insertionsort.Path))
	t.Log("Sorting time", time.Since(now))
}
