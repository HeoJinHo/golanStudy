package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// 파일 읽기
	file, err := os.Open("/Users/heojinho/Desktop/projects/GoProjects/firstProject/src/section11/sample.csv")
	errChaek43(err)

	// CSV Reader Created
	//rr := csv.NewReader(file)

	// 버퍼 권장
	rr := csv.NewReader(bufio.NewReader(file))

	// 내용 읽이
	// 1개의 row 별로 읽기
	row1, err1 := rr.Read()
	row2, err2 := rr.Read()
	errChaek43(err1)
	errChaek43(err2)

	fmt.Println("CSV Read Example")
	fmt.Println(row1[0], row1[1])
	fmt.Println(row2[0], row2[1])
	fmt.Println(len(row1))
	//fmt.Println(row)
	fmt.Println("==========================================================")

	// CSV 한번에 읽어오기
	rows, err := rr.ReadAll()
	errChaek43(err)
	fmt.Println("CSV ReadAll Example")
	// 가져온 데이터는 2차원 배열임
	fmt.Println(rows[5][1])

	fmt.Println(len(rows))

	for _, v := range rows {
		if v[0] == "Asia" {
			fmt.Println("This is ", v[0], " : ", v[1])
		}
	}

	// 리소스 해제
	defer file.Close()

}

// error check1
func errChaek43(e error) {
	if e != nil {
		panic(e)
	}
}

// error check2
func errChaek44(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}
