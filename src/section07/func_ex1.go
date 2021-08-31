package main

import "fmt"

func main() {

	a := []int{5, 6, 7, 8, 9, 10}

	m := multiply(a...)
	fmt.Println(m)
}

func multiply(a ...int) int {
	tot := 0

	for _, v := range a {
		tot += v
	}

	return tot
}
