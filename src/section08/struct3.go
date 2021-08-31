package main

import "fmt"

func main() {

	// Ex1
	c1 := Cars{"red", "22d"}

	c2 := new(Cars)
	c2.color, c2.name = "white", "sonata"

	c3 := &Cars{}

	c4 := &Cars{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println()

}

type Cars struct {
	color string
	name  string
}
