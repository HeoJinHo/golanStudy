package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 고루틴 동기화 기초

func main() {

	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, (기타언어 : notify, notifyAll)
	// Wait, Signal, Broadcas

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	// 비동기 버퍼 채널
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Wating : ", n)
			// 고루틴 멈춤
			condition.Wait()
			fmt.Println("Wating End")
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("received : ", <-c)
	}

	//for i := 0; i < 5; i++ {
	//	mutex.Lock()
	//	fmt.Println("Wake Goroutine(Signal) : ", i)
	//	// 한개씩 깨움, 모든 고루틴 끝나고
	//	condition.Signal()
	//	mutex.Unlock()
	//}

	// 뮤텍스 한번에 깨우기
	mutex.Lock()
	fmt.Println("Wake Goroutine(Boardcast) : ")
	condition.Broadcast()
	mutex.Unlock()
	time.Sleep(2 * time.Second)
}
