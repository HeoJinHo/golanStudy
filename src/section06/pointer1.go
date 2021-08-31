package main

import "fmt"

// 자료형 : 포인터(1)
func main() {

	// Pointer
	// Go : 포인터 지원(C)
	// 변수의 지역성, 연속된 메모리의 참조 ...., 힙, 스택
	// 주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	// *(아스테이크) 로 사용
	// nil 로 초기화 (nil == 0)

	// Ex1

	var a *int

	var b *int = new(int)

	fmt.Println(a)
	fmt.Println(b)

	i := 7

	a = &i
	b = &i

	fmt.Println()
	fmt.Println(a, &i, &a)
	// 역참조값
	fmt.Println(*a)
	fmt.Println(b)

}
