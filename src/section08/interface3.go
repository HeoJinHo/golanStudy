package main

import "fmt"

func main() {

	// 덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	//Go 의 덕타이핑 중요한 특징 : 오리처럼 걷고, 소리내고, 에엄 등 행동이 같으면 오리라고 볼 수 있다. -> 의미

	dog1 := Dog2{"poll", 10}
	cat1 := Cat2{"bob", 5}

	// Dog start
	act(dog1)

	// Cat start
	act(cat1)

}

// 구조체
type Dog2 struct {
	name   string
	weight int
}

type Cat2 struct {
	name   string
	weight int
}

// 동물의 행동 인터페이스
type Behavior2 interface {
	bite2()
	sounds2()
	run2()
}

// 인터페이스를 파라미터로 받는다.
func act(animal Behavior2) {
	animal.bite2()
	animal.sounds2()
	animal.run2()
}

// 구조체 Dog 메소드 구현
func (d Dog2) bite2() {
	fmt.Println(d.name, "Dog is bites!")
}

func (d Dog2) sounds2() {
	fmt.Println(d.name, "Dog is barks!")
}

func (d Dog2) run2() {
	fmt.Println(d.name, "Dog is running!")
}

// 구조체 Cat 메소드 구현
func (d Cat2) bite2() {
	fmt.Println(d.name, "Cat is 할퀴다!")
}

func (d Cat2) sounds2() {
	fmt.Println(d.name, "Cat is cries!")
}

func (d Cat2) run2() {
	fmt.Println(d.name, "Cat is running!")
}
