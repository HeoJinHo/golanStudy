package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 고루틴 고급(1)

func main() {

	// 고루틴 동기화 고급
	// Once : 한번만 실행 (주로 초기화에서 사용)
	// Do로 실행

	// Ex1

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Once 객체 생성
	once := new(sync.Once)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("go routine : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(2 * time.Second)

}

func onceTest() {
	// 이 부분에 한번 싱행 할 코드 작성
	fmt.Println("Once Test Excute!")
}
