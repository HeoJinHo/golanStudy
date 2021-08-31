package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	// 에러 처리 고급
	// Go errors 패키지 New 메소드 사용 -> 사용자 에러 처리 생성

	// Ex1
	if f, err := Power(7, 3); err != nil {
		fmt.Printf("Error Message : %s \n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	// Ex2
	if f, err := Power(0, 3); err != nil {
		fmt.Printf("Error Message : %s \n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}
}

// f 의 i 제곱 구하는 함수
func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용 불가능 합니다.")
	}

	// 제곱 nil 반환
	return math.Pow(f, i), nil
}
