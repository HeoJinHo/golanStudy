package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 고루틴
	// 클로저 사용 예제
	runtime.GOMAXPROCS(2)

	s := "Goroutine Closure : "

	// 반복문 클로저는 일반적으로 즉시 실행, 그러나 고루틴 클러저는 가장 나중에 실행(반복분 종료 후)
	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
	}

	time.Sleep(5 * time.Second)

}
