package main

import (
	"fmt"
	"strings"
)

func main() {

	// 예제1 (결합 : 일반연산)
	str1 := "Hello"
	str2 := "World"
	fmt.Println(str1 + ", " + str2)

	// 예제2 (결합 : Join)

	// string slice
	var strSet []string
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, ""))

}
