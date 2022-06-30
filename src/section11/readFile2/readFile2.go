// Read File(2)
package main

import (
	"bufio"
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
	// Read File - CSV file 읽기 예제

	// 파일 생성
	file, err := os.Open("../sample.csv")
	checkErr1(err)

	// 리소스 해제
	defer file.Close()

	// CSV Reader 생성
	// rr := csv.NewReader(file)	// 버퍼 사용 안함
	rr := csv.NewReader(bufio.NewReader(file)) // 권장 - 버퍼 사용

	// CSV 내용 읽기(1) - 1개 Row 단위 읽기
	row1, err := rr.Read() // 1개 Row 단위 읽기
	checkErr1(err)
	row2, err := rr.Read() // 1개 Row 단위 읽기
	checkErr1(err)
	checkErr1(err)
	fmt.Println("CSV Read Example")
	// fmt.Println(row)
	fmt.Println(row1[0], row1[1], row1[7], row1[1:5])
	fmt.Println(row2[0], row2[1], row2[7], row2[1:5])
	fmt.Println("==================")

	// CSV 내용 읽기(2) - 한 번에 읽기
	rows, err := rr.ReadAll() // 한 번에 읽기
	checkErr1(err)
	fmt.Println("CSV Read ALL Example")
	fmt.Println(rows)
	fmt.Println(rows[5])
	fmt.Println(rows[5][1])
	fmt.Println("==================")

	fileInfo, err := file.Stat()
	checkErr1(err)
	fmt.Println("파일 정보 출력(1) : ", fileInfo)
	fmt.Println("파일 이름(2) : ", fileInfo.Name())
	fmt.Println("파일 크기(3) : ", fileInfo.Size())
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime())
	fmt.Println("==================")

	for i, row := range rows {
		// fmt.Println(i, row)
		for j := range row {
			fmt.Printf("%20s ", rows[i][j])
		}
		fmt.Println()
	}

}
