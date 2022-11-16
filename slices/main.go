package main

import "fmt"

func main()  {
	s:=make([]string,3)
	fmt.Println(s)

	s[0]="a"
	s[1]="b"
	s[2]="c"

	fmt.Println(s)
	fmt.Println(s[1])

	fmt.Println(len(s))

	s=append(s, "d")
	s=append(s, "e","f")
	fmt.Println(s)

	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println(c)

	l:=s[2:5]
	fmt.Println(l)

	l=s[:5]
	fmt.Println(l)

	l=s[2:]
	fmt.Println(l)

	t:=[]string{"g","h","i"}
	fmt.Println(t)

	
	
}