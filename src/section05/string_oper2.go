package main

import "fmt"

func main() {

	// 예제2 비교

	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex1 : ", str1 == str2)
	fmt.Println("ex2 : ", str1 != str2)
	// GO에서 문자열 비교는 -> 아스키 코드에 대한 사전식 비교
	fmt.Println("ex3 : ", str1 > str2)
	fmt.Println("ex4 : ", str1 < str2)

}
