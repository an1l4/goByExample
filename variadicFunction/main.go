package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	for _, vals := range nums {
		total = total + vals
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

}
