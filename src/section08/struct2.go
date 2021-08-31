package main

import "fmt"

func main() {

	// 선언 방법1
	var kim *Accounts = new(Accounts)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	// 선언 방법2
	hong := &Accounts{"245-902", 15000000, 0.04}

	// 선언 방법3
	lee := new(Accounts)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("kim : ", kim)
	fmt.Println("hong : ", hong)
	fmt.Println("lee : ", lee)

	fmt.Printf("ex2 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)

	fmt.Println()

	// Ex1
	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(hong.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))

}

type Accounts struct {
	// 계좌번호
	number string

	// 잔액
	balance float64

	// 이자
	interest float64
}

func (a Accounts) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}
