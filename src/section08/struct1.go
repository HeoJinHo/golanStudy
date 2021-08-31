package main

import "fmt"

func main() {

	// Ex1

	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-902", balance: 12000000}
	park := Account{number: "245-903", interest: 0.025}
	cho := Account{interest: 0.025}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)

	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))

}

type Account struct {
	// 계좌번호
	number string

	// 잔액
	balance float64

	// 이자
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}
