package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, value := range nums {
		sum = sum + value
	}
	fmt.Println(sum)

	for index, value := range nums {
		if value == 3 {
			fmt.Println(index)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, ":", v)
	}

	for k := range kvs {
		fmt.Println("key is ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
