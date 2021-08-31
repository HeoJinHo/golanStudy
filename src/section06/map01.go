package main

import "fmt"

func main() {

	// Map
	// 맵 : 해시테이블, 딕셔너리(파이썬),  key-Value로 자료 저장
	// 레퍼런스 타입(참조 값 전달)
	// 비교 연산자 사용 불가능(참조 타입이여서)
	// 특징 : 참조타입(Key)로 사용 불가능, 값(Value)으로 모든 타입 사용가능
	// make  함수 및 축약(리터럴)로 초기화 가능
	// 순서 없음

	// Ex1

	var map1 map[string]int = make(map[string]int)
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println(map1, map2, map3)

	// Ex2
	map4 := map[string]int{}
	map4["apple"] = 25
	map4["banna"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banna":  40,
		"orange": 33,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banna"] = 40
	map6["orange"] = 33

	fmt.Println()
	fmt.Println(map4)
	fmt.Println(map5)
	fmt.Println(map6["orange"])
	fmt.Println(map6["apple"])

}
