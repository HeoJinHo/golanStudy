package main

import (
	"fmt"
	"os"
)

func main() {

	fileOpen("indefindisdls")

	fmt.Println("End Main!")

}

func fileOpen(filename string) {
	// 패닉이 발생되면 실행
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}

	defer f.Close()

}
