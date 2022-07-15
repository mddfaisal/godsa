package utils

import (
	"encoding/json"
	"io/ioutil"
)

func LongInputArr(path string) []int {
	var arr []int
	data, _ := ioutil.ReadFile(path)
	json.Unmarshal(data, &arr)
	return arr
}

func IsSameArr(a []int, b []int) bool {
	isSame := false
	if len(a) == len(b) {
		for i, key := range a {
			if key != b[i] {
				isSame = false
				break
			} else {
				isSame = true
			}
		}
	}
	return isSame
}
