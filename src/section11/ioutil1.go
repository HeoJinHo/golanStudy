package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	///Users/heojinho/Desktop/projects/GoProjects/firstProject/src/section11/sample.txt
	// 파일 읽기, 쓰기 -> oiutil 파키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// 아래 메소드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil
	// 읽기 : 4, 쓰기 : 2, 실행 : 1
	// 소유자, 그룹, 기타 사용자 순서(777)
	filename := "io_sample.txt"
	filename3 := "/Users/heojinho/Desktop/projects/GoProjects/firstProject/src/section11/sample.txt"

	// 파일 쓰기
	s := "Hello GoLang!\nFile Write Test! \n"
	err1 := ioutil.WriteFile(filename, []byte(s), os.FileMode(0644))

	errCheck(err1)

	data, err2 := ioutil.ReadFile(filename3)

	errCheck(err2)

	fmt.Println("===========================================")
	fmt.Println(string(data))
	fmt.Println("===========================================")

}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}
