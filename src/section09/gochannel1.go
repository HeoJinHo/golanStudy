package main

import (
	"fmt"
	"time"
)

// 채널(Channel) 기초(1)

func main() {

	// 채널
	// 고루틴 간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용
	// 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	// 데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고받을 수 있음)
	// interface{} 전달을 통해서 자료형 상관없이 전송 및 수신가능
	// 값을 전달 (복사 후 : bool, int등), 포인터(슬라이스, 맵) 등을 전달시에는 주의!
	// 멀티프로세싱 처리에서 교착상태(경쟁) 데드락 주의!
	// <-, -> 사용( 채널 <- 데이터 : 송신) /  ( 변수 <- 채널 : 수신)

	fmt.Println("Main : S -----> ", time.Now())

	//var c chan int
	//c = make(chan int)

	// int 형 체널 생성
	v := make(chan int)

	go work1(v)
	go work2(v)

	<-v
	<-v

	fmt.Println("Main : E -----> ", time.Now())

}

func work1(v chan int) {
	fmt.Println("Work1 : S ----> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ----> ", time.Now())
	// 1을 체널을 통해 송신
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : S ----> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ----> ", time.Now())
	// 1을 체널을 통해 송신
	v <- 1
}
