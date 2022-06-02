package main

import "fmt"

func main() {
	a := [6]int{1, 8, 6, 3, 5, 7}

	// append
	b := [7]int{}
	a = b

	fmt.Println(a)
}
