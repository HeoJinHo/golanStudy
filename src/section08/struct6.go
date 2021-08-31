package main

import (
	"fmt"
)

func main() {

	car1 := Carss2{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	fmt.Println("ex1 : ", car1.name, car1.color, car1.company)
	fmt.Printf("ex2 : %#v", car1.detail)

	// Ex2
	// 내부 spec 구조체 필드 값 출력
	fmt.Println("ex2 : ", car1.detail.length, car1.detail.height, car1.detail.width)
}

// 대문자로 선언시 패키지 외부에서 사용가능
type Carss2 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

// 첫글짜 소문자 선언시 패키지 내부에서만 사용가능
type spec struct {
	length int "저장"
	height int "전고"
	width  int "전축"
}
