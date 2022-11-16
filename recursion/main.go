package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter value")
	fmt.Scan(&n)
	fmt.Println(fact(n))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(n))

}

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)

}
