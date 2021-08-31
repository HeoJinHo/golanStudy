package main

import "fmt"

func main() {

	// 배열
	// 배열은 용량, 길이가 항상 같다.
	// 배열 vs 슬라이스 차이점 중요
	// 길이 고정 vs 길이 가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달
	// 전체 비교연산자 사용 가능 vs 비교 연산자 사용불가
	// 결론 : 대부분 슬라이스 사용한다. 거의 배열보다 슬라이스를 사용함

	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 갯수

	// Ex1

	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	// 배열길이 자동 증감 -> 들어오는 갯수만큼 길이 증감
	arr6 := [...]int{1, 2, 3, 4, 5}
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
	fmt.Println(arr7)
}
