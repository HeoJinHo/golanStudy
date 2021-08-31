package main

import "fmt"

func main() {

	//배열 순회

	// Ex1
	arr1 := [5]int{1, 10, 100, 1000, 1000}

	// len
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, arr1[i])
	}

	fmt.Println()

	// Ex2
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	fmt.Println()

	for _, v := range arr2 {
		fmt.Println(v)
	}

	// index 생략2

	for v := range arr2 {
		fmt.Println(v)
	}

}
