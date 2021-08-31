package main

import (
	"fmt"
	"time"
)

// 고루틴(Goroutine)(2)

func main() {

	//GoRoutine

	exe11("t1")

	// Ex1
	fmt.Println("Main Routine Start : ", time.Now())

	go exe11("t2")
	go exe11("t3")

	time.Sleep(4 * time.Second)

}

func exe11(name string) {
	fmt.Println(name, " start!", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>", i)
	}
	fmt.Println(name, " end!", time.Now())
}
