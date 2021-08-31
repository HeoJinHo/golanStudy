package main

// 사용자 패키지 설치 및 활용 예제

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {

	// 외부 저장소 패키지 설치
	// 2가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 1. go get 패키지 주소 설치 -> 선언

	// 선언 후 go get 설치 예제(엑셀 파일 읽기)
	xfile := "/Users/heojinho/Desktop/projects/GoProjects/firstProject/src/section12/sample.xlsx"

	xlFile, err := xlsx.OpenFile(xfile)
	errCheckP(err)

	for _, sheet := range xlFile.Sheets {
		sheet.ForEachRow(rowVisitor)
	}

}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s\t", value)

	}

	return err
}

func errCheckP(e error) {
	if e != nil {
		panic(e)
	}
}
