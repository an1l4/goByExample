package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("foo", "1")
	fmt.Println("Foo", os.Getenv("foo"))
	fmt.Println("BAR", os.Getenv("bar"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
