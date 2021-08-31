package main

import "fmt"
import "reflect"

// 인터페이스 고급(2)

func main() {

	// 타입 변환 (Type Assertion)
	// 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// 인터페이스.(타입) -> 형변환
	// interfaceVal.(type)

	var a interface{} = 15
	b := a
	c := a.(int)
	// 원본의 데이터 타입만으로만 돌아갈수 있음
	//d := a.(float64)

	fmt.Println(a, b, c)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

	// Ex2
	// 저장된 실제 타입 검사
	if v, ok := a.(float64); ok {
		fmt.Println("ex2 : ", v, ok)
	} else {
		fmt.Println("ex2 : ", ok)
	}

}
