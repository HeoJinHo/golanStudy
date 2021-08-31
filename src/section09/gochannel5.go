package main

// 채널(Channel) 기초(5)

import (
	"fmt"
)

func main() {

	// 채널(Channel)
	// Close : 채널 닫기, 주의 -> 닫힌 채널에 값 전송 시 패닉(예외) 발생
	// Range : 채널안에서 값을 꺼낸다. (순회), 채널 닫아야(Close) 반복문 종료 -> 채널이 열려 있고 값 전송하지 않으면 계속 대기!

	// Ex1
	ch := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
			fmt.Println("Chennel Open!")
		}
		// 5회 채널에 값 전송 후 채널 닫기
		close(ch)
	}()

	fmt.Println("Channel Close Start!")
	// 채널이 끝날때까지 대기
	for i := range ch {
		fmt.Println("ex1 : ", i)
	}

}
