package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res = plusPlus(2, 1, 3)
	fmt.Println(res)

}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
