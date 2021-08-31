package main

import "fmt"

func main() {

	//예제 1
	i := 7
	p := &i

	fmt.Println(&i, i, p, *p)

	*p++
	fmt.Println()
	fmt.Println(&i, i, p, *p)

	i = 77
	fmt.Println()
	fmt.Println(&i, i, p, *p)
}
