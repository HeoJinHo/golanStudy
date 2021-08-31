package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str0 string = "c:\\firstProject\\src"
	var str1 string = `c:\firstProject\src`

	fmt.Println(str1)
	fmt.Println(str0)

	var str3 = "Hello, World"
	var str4 = "안녕하세요."
	var str5 = "\ud55c\uae00"

	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)

	// 문자열의 바이트 length를 가져옴
	fmt.Println(len(str3))
	fmt.Println(len(str4))

	// 이방법을 선호함
	// 실제로 길이를 가져옴
	fmt.Println(utf8.RuneCountInString(str4))

	fmt.Println(len([]rune(str4)))

}
