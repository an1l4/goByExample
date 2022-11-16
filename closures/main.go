package main

import "fmt"

func main() {
	intNext := intSeq()
	fmt.Println(intNext())
	fmt.Println(intNext())
	fmt.Println(intNext())

	newInts := intSeq()
	fmt.Println(newInts())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
