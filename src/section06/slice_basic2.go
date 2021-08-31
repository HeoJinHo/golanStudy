package main

import "fmt"

func main() {

	// 슬라이스 참조타입 증명

	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1

	arr2[0] = 7

	fmt.Println(arr1, arr2)

	// slice

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1

	slice2[0] = 7

	fmt.Println(slice1, slice2)

	//에외 사항
	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	fmt.Println("ex3 : ", slice3[49])

	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}
}
