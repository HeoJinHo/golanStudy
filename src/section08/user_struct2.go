package main

import "fmt"

func main() {

	// 기본 자료형 사용자 정의 타입

	// Ex1

	a := cnt(15)

	var b cnt = 15

	fmt.Println(a, b)

	//testConverT(b)

	// 사용자 정의 타입 <-> 기본타입 : 매개변수 전달시 형변환 해서 전달해야한다
	testConverT(int(b))

}

type cnt int

func testConverT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex4 : (Custom Type) : ", i)
}
