package main

import (
	"fmt"
	"sync"
)

// 고루틴 고급(2)

func main() {

	// 고루틴 동기화 고급
	// 대기 그룹
	// 실행 흐름 변경(고루틴이 종료 될 때 까지 대기 가능)
	// WaitGroup : Add(고루틴 추가), Done(작업 졸료 알림), Wait(고루틴 종료시 까지 대기)
	// Add로 추 된 고루틴 갯수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다.(Add == Done)

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Wait Group : ", n)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같아야한다

	wg.Wait()
	fmt.Println("WaitGroup End!")

}
