package main

import (
	"fmt"
	"sort"
)

func main()  {
	strings:=[]string{"c","a","b"}
	sort.Strings(strings)
	fmt.Println(strings)

	ints:=[]int{7,2,5,8,1,4}
	sort.Ints(ints)
	fmt.Println(ints)

	s:=sort.IntsAreSorted(ints)
	fmt.Println(s)
	
}