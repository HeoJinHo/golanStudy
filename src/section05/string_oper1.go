package main

import "fmt"

func main() {

	// 문자열 연산
	// 문자열 추출, 비교, 조합(결합)

	// 예제1
	var str1 = "Golang"
	var str2 = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0])
	fmt.Println("ex2 : ", str2[3:], str2[0])

	fmt.Println("ex3 : ", str2[:4])
	fmt.Println("ex3 : ", str1[1:3])
}
