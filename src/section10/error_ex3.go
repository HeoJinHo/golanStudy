package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {

	// 예외(에러) 처리 구조체
	// error 타입이 아닌 경우 에러 처리 방법
	// 에러 메소드를 구현해서 사용자 정의 에러 처리 예제
	// 구조체를 사용해서 세부적인 에러처리

	// Ex1
	v, err := Power2(10, 3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("ex1 : ", v)

	// Ex2
	t, err := Power2(0, 3)
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.(*PowError).time)
		//fmt.Println(err.(*PowError).value)
		//fmt.Println(err.(*PowError).message)
	}

	fmt.Println("ex2 : ", t)

}

type PowError struct {
	// Error Time
	time time.Time
	//parameter
	value float64
	// Error Message
	message string
}

func (e *PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

func Power2(f, i float64) (float64, error) {
	if f == 0 {
		return 0, &PowError{time.Now(), f, "0은 입력 할 수 없습니다."}
	}

	if math.IsNaN(f) {
		return 0, &PowError{time.Now(), f, "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, &PowError{time.Now(), f, "숫자가 아닙니다."}
	}

	return math.Pow(f, i), nil
}
