package main

import "fmt"

func main() {

	// Channel
	// 채널 또한 함수의 반환 값으로 사용 가능

	c := sum(100)

	fmt.Println("ex1 : ", <-c)
}

// 반환형 채널임
func sum(cnt int) <-chan int {
	sum := 0

	tot := make(chan int)

	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}
