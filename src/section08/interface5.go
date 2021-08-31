package main

import "fmt"

func main() {

	// 인터페이스 활용(비어있는 인터페이스)
	// 함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능)
	// 모든 타입 지정 가능

	dog1 := Dog5{"poll", 10}
	cat1 := Cat5{"bob", 5}

	printValue(dog1)
	printValue(cat1)

	printValue(15)
	printValue("Animal")
	printValue(25.5)
	printValue([]Dog5{})
	printValue([5]Dog5{})
	printValue(true)

}

func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

// 구조체
type Dog5 struct {
	name   string
	weight int
}

type Cat5 struct {
	name   string
	weight int
}
