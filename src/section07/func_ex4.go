package main

import "fmt"

func main() {

	// function 고급
	// 익명 함수(Anonymous Function)
	// 즉시 실행 (선언과 동시에 실행)

	// Ex1

	func() {
		fmt.Println("ex1 : Anonymous Function!")
	}()

	msg := "Hello GoLang"

	func(msg string) {
		fmt.Println("ex2 : ", msg)
	}(msg)

	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20)

	// Ex4
	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println("ex4 : ", r)

}
