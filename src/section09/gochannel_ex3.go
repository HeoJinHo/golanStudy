package main

import "fmt"

func main() {

	// Channel
	// 채널 또한 함수의 반환 값으로 사용 가능

	// channel return
	c := receiveOnly2(100)

	// channel send return channel
	output := total(c)

	fmt.Println("ex1 : ", <-output)
}

func receiveOnly2(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a += i
		}
		tot <- a
	}()

	return tot
}
