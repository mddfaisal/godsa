package main

import (
	"arr_sorting/input"
	"fmt"
	"time"
)

var ip []int = input.LongInputIntArr()

func main() {
	insertionsort()
	bubblesort()
	mergesortalgo()
}

func bubblesort() {
	t := time.Now()
	arr := ip

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	fmt.Println("Bubble sort")
	fmt.Println("Elapsed: ", time.Since(t))
}

func insertionsort() {
	t := time.Now()
	arr := ip

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	fmt.Println(arr)
	fmt.Println("Insertion sort")
	fmt.Println("Elapsed: ", time.Since(t))
}

func mergesortalgo() {
	t := time.Now()
	arr := ip
	fmt.Println(mergesort(arr))
	fmt.Println("Merge sort")
	fmt.Println("Elapsed: ", time.Since(t))
}

func mergesort(arr []int) []int {
	size := len(arr)
	if size == 1 {
		return arr
	}
	middle := int(size / 2)
	left := make([]int, middle)
	right := make([]int, size-middle)

	for i := 0; i < size; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}
	return merge(mergesort(left), mergesort(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
