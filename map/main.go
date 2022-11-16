package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["k1"] = 7
	m["k2"] = 3

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(m)

	d, prs := m["k2"]
	fmt.Println(d, prs)

}
