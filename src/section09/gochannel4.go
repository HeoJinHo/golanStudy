package main

import (
	"fmt"
	"runtime"
)

// 채널(Channel) 기초(4)

func main() {

	// 채널(Channel)
	// Ex1 (비동기 : 버퍼 사용)
	// 버퍼 : 발신 -> 가득차면 대기, 비어있으면 다시 작동 , 수신 -> 비어있으면 대기, 가득차 있으면 작동

	runtime.GOMAXPROCS(4)

	// 비동기 일 경우는 채널의 버퍼를 지정할 수 있음
	ch := make(chan bool, 4)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
