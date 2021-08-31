package main

import "fmt"

func main() {

	// Go Recover function
	// 에러 복구 가능
	// Panic 으로 발생한 에러를 복구 후 코드 재 싱행(프로그램 종료 되지 않는다.)
	// 정상 상태로 복구 하는 기능
	// 패닉에서 설정한 메세지를 받아올 수 있다.

	// Ex1

	runFunc2()

	fmt.Println("Heloo GoLang")

}

func runFunc2() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Massage : ", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", a[i])
	}

}
