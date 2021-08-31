package main

import "fmt"

func main() {

	kim := Acctouns4{"245-901", 10000000, 0.015}
	lee := Acctouns4{"245-902", 20000000, 0.025}

	lee.CalculateD4(1000000)
	kim.CalculateP4(1000000)

	fmt.Println("ex2 : ", int(lee.balance))
	fmt.Println("ex2 : ", int(kim.balance))

}

type Acctouns4 struct {
	number   string
	balance  float64
	interest float64
}

func (a *Acctouns4) CalculateD4(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a Acctouns4) CalculateP4(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}
