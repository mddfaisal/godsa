package twosum_test

import (
	"godsa/problems/arrays/twosum"
	"godsa/utils"
	"testing"
)

var (
	arr    = []int{2, 7, 11, 15}
	result = []int{0, 1}
	target = 9
)

func TestTwoSum(t *testing.T) {
	if !utils.IsSameArr(twosum.TwoSum(arr, target), result) {
		t.Error("Arrays are not same.")
	}
}
