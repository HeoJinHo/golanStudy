package main

import "fmt"

func main() {

	// 구조체 임베디드 메소드 오버라이딩 패턴

	// Ex1

	ep1 := Employee2{"kim", 2000000, 150000}
	ep2 := Employee2{"park", 1500000, 200000}

	//임원

	ex := Executives2{
		Employee2{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex2 : ", int(ex.Calculate()))
	fmt.Println("ex2 : ", int(ex.Employee2.Calculate()+ex.specialBonus))
}

type Employee2 struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee2) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives2) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives2 struct {
	// is a 관계
	Employee2
	specialBonus float64
}
