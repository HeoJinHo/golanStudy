package main

import (
	"fmt"
	"os"
)

func main() {

	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의(오류 메세지 확인)
	// 예외 처리 중요
	// 탐색 Seek 중요
	// https://golang.org/pkg/os

	// 파일 읽기 예제
	// 파일 열기

	file, err := os.Open("/Users/heojinho/Desktop/projects/GoProjects/firstProject/src/section11/sample.txt")

	errChaek31(err)

	// Read Ex1
	// 파일 사이즈 확인 위해 정보 획득
	fileInfo, err := file.Stat()
	errChaek32(err)
	// 슬라이스에 읽은 내용 담기 / 사이즈는 효율을 높이기 위해
	fd1 := make([]byte, fileInfo.Size())
	ct1, err := file.Read(fd1)
	errChaek31(err)

	fmt.Println("파일 정보 출력(1) : ", fileInfo, "\n")
	fmt.Println("파일 이름 출력(2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기 출력(3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정시간 출력(4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n", ct1)

	// 타입변환 없을경우
	fmt.Println(fd1)
	fmt.Println()
	// 타입변환 한경우
	fmt.Println(string(fd1))

	fmt.Println("===================================================================")

	// 읽기 예제2(탐색 : Seek(offset))
	// whence -> 0 : 처음 위치, 1: 현재 위치, 2: 마지막 위치
	o1, err := file.Seek(20, 0)

	errChaek31(err)
	fd2 := make([]byte, 20)

	ct2, err := file.Read(fd2)
	errChaek31(err)

	fmt.Printf("읽기 작업(1) 완료 (%d bytes)(%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")

	// 읽기 예제3
	o2, err := file.Seek(0, 0)
	errChaek31(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8)
	errChaek31(err)
	fmt.Printf("읽기 작업(3) 완료 (%d bytes)(%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")
	fmt.Println("===================================================================")

}

// error check1
func errChaek31(e error) {
	if e != nil {
		panic(e)
	}
}

// error check2
func errChaek32(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}
