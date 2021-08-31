package main

import "fmt"

func main() {

	// 클로저 (Closure)
	// 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수안에서 함수 선언 및 정의 가능 -> 이 때 외부 함수에 선언 된 변수에 접근 사능한 함수
	// 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 남발 -> 객체 들이 메모리에 존재하므로, -> 메모리부족, 오버플로우 현상 발생가능성이 있음
	// 클로저 정확하게 이해하고 사용해야함

	// 익명함수 클로저 사용안함
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println("r1 : ", r1)

	// closure used

	m, n := test1(), test1()

	sum := func(c int) int { // 익명함수 변수 할당
		return m + n + c // 지역변수 소멸되지 않음
	}

	r2 := sum(10)

	fmt.Println("ex2: ", r2)

}

func test1() int {
	return 10
}
