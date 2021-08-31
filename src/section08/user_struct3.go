package main

import "fmt"

func main() {

	var orderPrice totCost

	orderPrice = func(cnt int, price int) int {
		return (cnt * price) + 100000
	}

	describe(3, 5000, orderPrice)

}

// function usert struct
type totCost func(int, int) int

// 함수 재정의
func describe(cnt, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, orderPrice : %d", cnt, price, fn(cnt, price))
}
