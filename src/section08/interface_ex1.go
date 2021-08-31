package main

import "fmt"

// 인터페이스 고급(1)

func main() {

	// 비어있는 인터페이스 활용(빈 인터페이스)
	// 함수 매개변수, 리턴 값, 구조체 필드 등으로 사용 가능 -> 어떤 타입으로도 변환 가능
	// 모든 타입을 나타내기 위해 비어있는 인터페이스 사용
	// 동적타입 으로 생각하면 쉽다.(타 언어의 Object 타입)

	printContents(15)
	printContents("Animal")
	printContents(25.5)
	printContents(true)

}

func printContents(v interface{}) {
	// 원본 타입
	fmt.Printf("Type : (%T)", v)
	fmt.Println("ex1 : ", v)

}
