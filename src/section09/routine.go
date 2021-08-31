package main

import (
	"fmt"
	"time"
)

// 고루틴(Goroutine)

func main() {

	// 고루틴
	// 타 언어의 쓰레드와 비슷한 기능
	// 생성 방법 매우 간단, 리소스 매우 적게 사용
	// 즉, 수많은 고루틴 동시 생성 실행 가능
	// 비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통한 통신 가능
	// 공우메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글루틴(메인메소드)에 비해 항상 빠른 처리 결과는 아니다.

	// 멀티 스레드의 장점과 단점
	// 장점 : 개발이 잘되면 응답성 향상, 자원 공유를 효율적으로 활용 사용, 작업이 분리되어 코드 간결
	// 단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스에 사이드이펙트,  성능 저하, 동기화 코딩 반드시 숙지, 데드락
	// 교착상태 등등 단점이 더 많음

	// 일반적인 실행
	//exe1()
	//exe2()
	//exe3()

	// 가장 먼저 실행(일반적인 실행 흐름)
	// 메인 함수가 종료가 되면 그뒤에 고루틴이 있어도 같이 종료됨
	exe1()

	fmt.Println("Main Rouine Start!", time.Now())

	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Rouine End!", time.Now())

}

func exe1() {
	fmt.Println("ex1 function start!", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ex1 function end!", time.Now())
}

func exe2() {
	fmt.Println("ex2 function start!", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ex2 function end!", time.Now())
}

func exe3() {
	fmt.Println("ex3 function start!", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("ex3 function end!", time.Now())
}
