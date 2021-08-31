package main

import "fmt"

func main() {

	// 덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	//Go 의 덕타이핑 중요한 특징 : 오리처럼 걷고, 소리내고, 에엄 등 행동이 같으면 오리라고 볼 수 있다. -> 의미

	dog1 := Dog3{"poll", 10}
	cat1 := Cat3{"bob", 5}

	// Dog start
	act3(dog1)

	// Cat start
	act3(cat1)

}

// 구조체
type Dog3 struct {
	name   string
	weight int
}

type Cat3 struct {
	name   string
	weight int
}

// 동물의 행동 인터페이스
type Behavior3 interface {
	run3()
}

// 인터페이스를 파라미터로 받는다.
func act3(animal interface{ run3() }) {
	animal.run3()
}

func (d Dog3) run3() {
	fmt.Println(d.name, "Dog is running!")
}

// 구조체 Cat 메소드 구현
func (d Cat3) run3() {
	fmt.Println(d.name, "Cat is running!")
}
