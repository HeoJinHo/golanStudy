package main

import "fmt"

func main() {

	bs1 := shoppingBasket{3, 5000}
	//bs2 := shoppingBasket{5000, 1}

	fmt.Println("ex1(totPrice) : ", bs1.purchase())

	bs1.rePurchaseP(7, 5000)

	// 원본값 수정 되었음
	fmt.Println("ex1(rePurchaseP) : ", bs1.purchase())

	bs1.rePurchaseD(10, 0)
	// 원본값 수정 안되었음
	fmt.Println("ex1(rePurchaseP) : ", bs1.purchase())

}

type shoppingBasket struct {
	cnt, price int
}

// 결제 함수
func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 참조형을 받기 때문에 원본값 수정
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본수정 안됨 값 전달 형식
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}
