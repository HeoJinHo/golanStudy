package main

// 고루틴(Goroutine)(3)

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	// Ex1
	// 멀티 코어 다중 CPU 최대한 활용
	runtime.GOMAXPROCS(runtime.NumCPU())
	// GOMAXPROCS(0)을 하면 위에서 설정한 CPU 몇개 사용한지 출력해줌
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0))

	fmt.Println("Main Routine Start : ", time.Now())

	for i := 0; i < 100; i++ {
		go exe33(i)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Main Routine End : ", time.Now())

}

func exe33(name int) {
	//100 미만의 랜덤 정수
	r := rand.Intn(100)
	fmt.Println(name, "Start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>>> ", r, " : ", i)
	}
	fmt.Println(name, "End : ", time.Now())
}
