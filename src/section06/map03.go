package main

import "fmt"

func main() {

	// Map
	// 맵 조회 및 순회(Iterator)

	map1 := map[string]string{
		"daum":   "daum.net",
		"naver":  "naver.com",
		"google": "goole.com",
		"home1":  "test1.com",
	}

	fmt.Println(map1)

	// 추가 java == map.put("home2", "test2.com")
	map1["home2"] = "test2.com"

	fmt.Println(map1)

	//수정
	map1["home1"] = "test11.com"

	fmt.Println(map1)

	fmt.Println()

	//삭제
	delete(map1, "home2")
	delete(map1, "google")
	fmt.Println(map1)

}
