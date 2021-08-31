package main

import "fmt"

func main() {

	// Map
	// 맵 조회 및 순회(Iterator)

	map1 := map[string]string{
		"daum":   "daum.net",
		"naver":  "naver.com",
		"google": "goole.com",
	}
	fmt.Println(map1["google"])
	fmt.Println()

	// Ex 순서가 없음으로 랜덤
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println(v)
	}
}
