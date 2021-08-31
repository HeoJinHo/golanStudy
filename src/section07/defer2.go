package main

import (
	"fmt"
)

func main() {

	// Ex1

	sayHello("GoLang")

}

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("H2")
	}()
}
