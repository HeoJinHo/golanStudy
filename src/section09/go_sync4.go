package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 고루틴 동기화 기초

func main() {

	// 뮤텍스 : 상호 배제 -> Tread(고루틴)들이 서로 RunningTime 에 서로 영향을 주지 않게 즉, 단독으로 실행되게 하는 기술

	// 동기화 사용하는 경우 예제
	// RWMutex : 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 X, 읽기 락, 쓰기 락 전부 방버(방지)
	// RMutex : 읽기 Lock -> 읽기 시도 중에 값 변경 방지 즉, 쓰기 락 방어(방지)

	// Ex1

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)

			// 쓰기 뮤텍스 잠금 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)

			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)

			// 읽기 뮤텍스 잠금 해제
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)

}
