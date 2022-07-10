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
