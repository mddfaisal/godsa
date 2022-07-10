package linearsearch_test

import (
	"searching-sorting/searching/linearsearch"
	"searching-sorting/utils"
	"testing"
	"time"
)

func TestLinearSearch(t *testing.T) {
	t.Log("Linear search...")
	arr := utils.LongInputArr(linearsearch.Path)
	now := time.Now()
	linearsearch.Search(arr, 999)
	t.Log("Linear search time: ", time.Since(now))
}
