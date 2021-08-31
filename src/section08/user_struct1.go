package main

import "fmt"

// 사용자 정의 타입(1)

func main() {

	// Go -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
	// 객체 지향 -> 클래스(속성, 기능(상태)) : 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그래밍
	// Go는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스형태의 코딩 가능
	// 객체지향의 기본 개념 -> Go에서 포함 하고 있다. -> 객체 지향 프로그래밍 언어
	// 상태, 메소드 분리해서 정의(결합성 없음)
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본타입(int, float, string ...), 함수
	// 구조체와 -> 메소드 연결을 통해서 타 언어의 클래스 형식 처럼 사용가능(객체지향)

	// Ex1

	bmw := Car{name: "520d", price: 50000000, color: "white", tax: 5000000}
	benz := Car{name: "220d", price: 60000000, color: "blue", tax: 6000000}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex2 : ", benz, &benz)

	fmt.Println("ex3 : ", Price(bmw))

	// 구조체랑 연결된 메소드
	fmt.Println("ex4 : ", bmw.Price2().price, bmw.Price2().tax)

}

// 구조체 <-> 메소스 바인딩
// 아래 메소드는 일반 메소드임
func Price(c Car) int64 {
	return c.price + c.tax
}

// 구조체랑 연결된 메소드
func (c *Car) Price2() Car {
	c.price = 1000
	c.tax = 10
	return *c
}

// 구조체 선언
type Car struct {
	name  string
	color string
	price int64
	tax   int64
}
