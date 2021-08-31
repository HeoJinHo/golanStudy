package main

import "fmt"

func main() {

	// 재귀함수(Recursion)

	// Ex1
	x := fact(7)

	fmt.Println(x)

	prtHello(10)
}

func prtHello(n int) {
	if n == 0 {
		return
	}

	fmt.Println("ex2: , (", n, ")", "hi!")

	prtHello(n - 1)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
