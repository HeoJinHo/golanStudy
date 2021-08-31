package main

import "fmt"

func main() {

	kim := Acctouns3{"245-901", 10000000, 0.015}
	lee := Acctouns3{"245-902", 20000000, 0.025}

	fmt.Println(kim)
	fmt.Println(lee)

	fmt.Println()

	CalculateD2(kim)
	CalculateP2(&lee)

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))

}

type Acctouns3 struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD2(a Acctouns3) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP2(a *Acctouns3) {
	a.balance = a.balance + (a.balance * a.interest)
}
