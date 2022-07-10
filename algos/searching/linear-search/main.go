package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func longinput() []int {
	var arr []int
	data, _ := ioutil.ReadFile("../../input/long.arr")
	json.Unmarshal(data, &arr)
	return arr
}

var (
	i = 139211814
)

func main() {
	linearsearch1()
}

func linearsearch1() {
	arr := longinput()
	t := time.Now()
	found := false

	for _, v := range arr {
		if v == i {
			found = true
			break
		}
	}
	fmt.Println("linearsearch1: ", time.Since(t), found)
}
