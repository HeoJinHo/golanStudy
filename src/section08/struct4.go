package main

import "fmt"

func main() {

	// 구조체 익명 선언 및 활용

	car1 := struct {
		name, color string
	}{"520d", "red"}

	fmt.Println(car1)
	fmt.Printf("ex1: %#v \n", car1)

	cars := []struct {
		name, color string
	}{
		{"520d", "red"},
		{"530d", "white"},
		{"528d", "blue"},
	}

	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v) \n", c.name, c.color, c)
	}

}
