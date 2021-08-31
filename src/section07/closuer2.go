package main

import "fmt"

func main() {

	cnt := increaseCnt()

	fmt.Println("cnt1 : ", cnt())
	fmt.Println("cnt1 : ", cnt())
	fmt.Println("cnt1 : ", cnt())
	fmt.Println("cnt1 : ", cnt())
	fmt.Println("cnt1 : ", cnt())

}

func increaseCnt() func() int {
	// 현재 함수의 지역변수(캡쳐됨)
	n := 0

	return func() int {
		n += 1
		return n
	}
}
