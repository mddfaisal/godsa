package main

import (
	"fmt"
	"time"
)

var (
	p float64 = 1
	f float64 = 1
)

func e_recursive(x float64, n float64) float64 {
	var r float64
	if n == 0 {
		return 1
	}
	r = e_recursive(x, n-1)
	p = p * x
	f = f * n
	return r + p/f
}

func e_iterative(x float64, n float64) float64 {
	var t, num, den float64
	t, num, den = 1, 1, 1
	for i := 1; i <= int(n); i++ {
		num = num * x
		den = den * float64(i)
		t = t + num/den
	}
	return t
}

func fact(n float64) float64 {
	if n == 0 || n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fact(n-1) * n
}

func ncr(n float64, r float64) float64 {
	var f float64 = 1
	for r >= 1 {
		f = f * (n / r)
		r--
		n--
	}
	return f
}

func main() {
	var n, x, r float64

	fmt.Print("Enter x: ")
	fmt.Scanf("%f", &x)

	fmt.Print("Enter n: ")
	fmt.Scanf("%f", &n)

	fmt.Print("Enter r: ")
	fmt.Scanf("%f", &r)

	t := time.Now()
	// fmt.Println("using recursion: ", e_recursive(x, n))
	// fmt.Println("time taken: ", time.Since(t))

	// t = time.Now()
	// fmt.Println("using iteration: ", e_iterative(x, n))
	// fmt.Println("time taken: ", time.Since(t))

	t = time.Now()
	fmt.Println("ncr using fact: ", fact(n)/(fact(r)*fact(n-r)))
	fmt.Println("time taken: ", time.Since(t))

	t = time.Now()
	fmt.Println("ncr using loop: ", ncr(n, r))
	fmt.Println("time taken: ", time.Since(t))

}
