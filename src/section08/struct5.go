package main

import (
	"fmt"
	"reflect"
)

func main() {

	tag := reflect.TypeOf(Carss{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

}

type Carss struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}
