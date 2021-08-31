package main

import (
	"fmt"
	"time"
)

// 채널(Channel) 기초(3)

func main() {

	// 채널(Channel)
	// Ex1 (동기 : 버퍼 미사용)

	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
