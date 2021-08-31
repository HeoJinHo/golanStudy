package main

import "fmt"

// 인터페이스 고급(3)

func main() {

	// 실제 타입 검사 switch 사용
	// 비어있는 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능

	// Ex1
	checkType(true)
	checkType(1)
	checkType(22.242)
	checkType(nil)
	checkType("Hello GoLang")

}

func checkType(arg interface{}) {
	// switch
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool -> ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int -> ", arg)
	case float64:
		fmt.Println("This is a float -> ", arg)
	case string:
		fmt.Println("This is a string -> ", arg)
	case nil:
		fmt.Println("This is a nil -> ", arg)
	default:
		fmt.Println("What is this type? -> ", arg)
	}

}
