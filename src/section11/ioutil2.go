package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	///Users/heojinho/Desktop/projects/GoProjects/firstProject/src/section11/sample.txt

	// 파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, Writer, Read 메소드를 사용 가능
	/*
		type Reader interface{
			Read(p []byte) (n int, err error)
		}
		type Writer interface{
			Write(p []byte) (n int, err error)
		}
	*/
	// 즉, bufio의 NewReader, NweWriter를 통해서 객체 반환 후 메소드 사용
	// bufio(Buffered io) 패키지
	// https://golang.org/pkg/bufio

	// 파일 열기
	// 두 번째 매개변수 확인
	// https://golang.org/pkg/os/#pkg-constants

	/*
		상태
		a ---> a
		b ---> ab
		c ---> abc
		d ---> abcd //버퍼 꽉참
		e ---> e 	--------> abdc
		f ---> ef 	--------> abdc
		g ---> eg 	--------> abdc
		h ---> egh 	--------> abdc
		i ---> i 	--------> abdcefgh
	*/

	// os.O_CREATE : 파일이 없으면 생성
	// os.O_RDWR : 일고 쓰는 권한 부여
	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck2(err)

	// bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file)
	wt.WriteString("1. Hello GoLang!\nFile Write Test!\n")
	wt.Write([]byte("2. Hello GoLang!\nFile Write Test!\n"))

	// 사용한 버퍼 정보 출력
	fmt.Printf("사용한 Buffer Size (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d bytes)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d bytes)\n", wt.Size())

	// 이때 버퍼 비우고 반영
	errCheck2(wt.Flush())

	//Reader 반환
	rt := bufio.NewReader(file)

	fi, err := file.Stat()
	errCheck2(err)

	b := make([]byte, fi.Size())
	fmt.Println("========================================")
	fmt.Println("파일 정보 출력(1) : ", fi, "\n")
	fmt.Println("파일 이름 출력(2) : ", fi.Name(), "\n")
	fmt.Println("파일 크기 출력(3) : ", fi.Size(), "\n")
	fmt.Println("파일 수정시간 출력(4) : ", fi.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n", b)
	fmt.Println("========================================")

	file.Seek(0, os.SEEK_SET)
	// 읽기
	data, _ := rt.Read(b)
	fmt.Printf("전체 버퍼 사이즈 : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)

	fmt.Println("========================================")
	fmt.Println(string(b))
	fmt.Println("========================================")

	defer file.Close()
}

func errCheck2(e error) {
	if e != nil {
		panic(e)
	}
}
