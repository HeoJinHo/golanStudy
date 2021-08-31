package main

import (
	"fmt"
	"os"
)
import "encoding/csv"

func main() {

	// 파일 쓰기 csv
	// CSV 파일쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식, 쓰기, 읽기 가능
	// bufio(버퍼IO) : 파일 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("test_write2.csv")

	errChaek11(err)

	// 리소스 해제
	defer file.Close()

	// CSV Writer Created
	wr := csv.NewWriter(file)

	// 큰파일 사용할때
	//wr := csv.NewWriter(bufio.NewWriter(file))

	// CSV 내용 쓰기
	wr.Write([]string{"kim", "4.5"})
	wr.Write([]string{"lee", "4.4"})
	wr.Write([]string{"park", "4.3"})
	wr.Write([]string{"jo", "4.2"})
	wr.Write([]string{"hong", "4.1"})

	// 버퍼 -> 파일로 쓰고 비우기
	wr.Flush()

	fi, err := file.Stat()

	errChaek11(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())

}

// error check1
func errChaek11(e error) {
	if e != nil {
		panic(e)
	}
}

// error check2
func errChaek22(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}
