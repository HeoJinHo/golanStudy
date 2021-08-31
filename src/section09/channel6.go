package main

import "fmt"

// 채널(Channel) 기초(6)

func main() {

	// Close : 채널 닫기

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 7777
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex2 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex3 : ", val3, ok3)

	// 채널 닫기
	close(ch)

	// 채널이 닫힌 상태여서 ok는 false
	val4, ok4 := <-ch
	fmt.Println("ex4 : ", val4, ok4)

}
