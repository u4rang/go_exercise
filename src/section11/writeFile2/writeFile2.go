// Write File(2)
package main

import (
	_ "bufio"
	"encoding/csv"
	_ "encoding/csv"
	"fmt"
	_ "fmt"
	"os"
	_ "os"
)

// Check Error 1.
func checkErr1(err error) {
	if err != nil {
		panic(err)
	}
}

// Check Error 2.
func checkErr2(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// Write File
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용권장

	// Create File
	file, err := os.Create("test.csv")

	// 리소스 해제
	defer file.Close()

	// csv writer 생성
	wr := csv.NewWriter(file)

	// wr := csv.NewWriter(bufio.NewWriter(file))	// 권장방식

	// csv 내용 쓰기
	wr.Write([]string{"kim", "4.8"})
	wr.Write([]string{"lee", "4.5"})
	wr.Write([]string{"park", "4.7"})
	wr.Write([]string{"choi", "4.9"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err := file.Stat()

	checkErr1(err)
	fmt.Printf("CSV 쓰기 작업 후 파일 크기 : %d bytes\n", fi.Size())
	fmt.Printf("CSV 파일 명 : %s\n", fi.Name())
	fmt.Printf("운영체제 파일 권한 : %s\n", fi.Mode())

}
