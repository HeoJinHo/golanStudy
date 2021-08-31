package main

import (
	"fmt"
	"os"
)

func main() {

	// 파일 쓰기
	// Create :  새 파일 작성 및 파일 열기
	// Close : 리소스 닫기
	// Write, WriteString, WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의()
	// 예외 처리 중요!
	// https://golang.org/pkg/os

	file, err := os.Create("test_write.txt")
	errChaek1(err)

	defer file.Close()

	// 파일 쓰기 예제1
	s1 := []byte{48, 49, 50, 51, 51}
	// 문자열 -> 바이트 슬라이스 형으로 변환 후 쓰기
	n1, err := file.Write([]byte(s1))
	errChaek2(err)

	fmt.Printf("Write Finsh (%d bytes)\n", n1)

	// 파일쓰기를 반영하는것 commit 같은 역할
	file.Sync()

	// 쓰기 예제2
	s2 := "Hello GoLang! \n File Write Test! - 1 \n"
	n2, err := file.WriteString(s2)
	fmt.Printf("Write Finsh (%d bytes)\n", n2)
	errChaek2(err)

	s3 := "Test WriteAt! - 2 \n"
	n3, err := file.WriteAt([]byte(s3), 70)

	errChaek1(err)

	fmt.Printf("Write Finsh (%d bytes)\n", n3)

	file.Sync()

	n4, err := file.WriteString("\n Hello Golang! \n File Write Test! - 3")

	errChaek1(err)

	fmt.Printf("Write Finsh (%d bytes)\n", n4)

}

// error check1
func errChaek1(e error) {
	if e != nil {
		panic(e)
	}
}

// error check2
func errChaek2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}
