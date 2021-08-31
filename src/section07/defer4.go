package main

import "fmt"

func main() {
	// Ex1
	a()
}

func a() {
	// defer 문이 있더라도 defer 문법 다음에 있는 함수만 마지막에 호출함
	defer end(start("b"))
	fmt.Println("in a")
}

func end(t string) {
	fmt.Println("End : ", t)
}

func start(msg string) string {
	fmt.Println("Start : ", msg)
	return msg
}
