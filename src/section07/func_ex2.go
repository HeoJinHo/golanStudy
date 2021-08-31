package main

import "fmt"

func main() {

	f := []func(int, int) int{multyply2, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println(a, b)

	// 변수에 할당
	var f1 func(int, int) int = multyply2
	f2 := sum

	fmt.Println(f1(10, 10), f2(10, 10))

	m := map[string]func(int, int) int{
		"mulFunc": multyply2,
		"sumFunc": sum,
	}

	fmt.Println(m["mulFunc"](10, 10), m["sumFunc"](10, 10))

}

func multyply2(x, y int) (r int) {
	r = x * y

	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}
