package utils

import (
	"encoding/json"
	"godsa/datatype"
	"io/ioutil"
)

func Interface2Val(i interface{}) interface{} {
	var val interface{}
	switch v := i.(type) {
	case int:
		val = v
	case *datatype.TNode:
		val = v.Key
	default:
		val = ""
	}
	return val
}

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
