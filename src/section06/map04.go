package main

import "fmt"

func main() {

	// 맵 조회 할 경우에 주의 할점

	map1 := map[string]int{
		"apple":  15,
		"banna":  115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	// 실질적으로 값은 두개를 리턴함 : 두번째 값은 key의 존재여부를 리턴함
	value3, ok := map1["kiwi"]

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3, ok)

	fmt.Println()
	if value, ok := map1["kiwi"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("kiwi is not exist")
	}

	fmt.Println()
	if value, ok := map1["banna"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("banna is not exist")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("banna is not exist")
	}

}
