package main

import "fmt"

func main() {

	// 구조체 생성자 패턴 예제
	kim := Acctouns2{"245-901", 10000000, 0.015}

	var lee *Acctouns2 = new(Acctouns2)
	lee.number = "245-902"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	park := NewAccount("245-903", 17000000, 0.04)

	fmt.Println("ex2 : ", park)

}

type Acctouns2 struct {
	number   string
	balance  float64
	interest float64
}

// 포인터 반환 아닌경우는 값 복사
// 아래 예제는 포인터를 반환함으로 값 참조하여 변환함(원본 데이터 바뀜)
func NewAccount(number string, balance, interest float64) *Acctouns2 {
	// 구조체 인스턴스를 생성 한 뒤 포인터를 반환
	return &Acctouns2{number, balance, interest}
}
