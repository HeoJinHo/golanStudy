package main

import "fmt"

// 채널(Channel) 기초(2)

func main() {

	// 채널(Channel)

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 데이터 수신(동기) : 체널에서 값 수신될때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("result1 : ", result1)
	fmt.Println("result2 : ", result2)
	fmt.Println("result3 : ", result3)

}

func rangeSum(rg int, c chan int) {
	sum := 0
	for i := 1; i < rg; i++ {
		sum += i
	}

	c <- sum
}
