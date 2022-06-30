// Read File(1)
package main

import (
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
	// Read File
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요
	// 탐색 Seek 중요
	// https://golang.org/pkg/os

	// 파일 읽기 예제
	// 파일 읽기
	file, err := os.Open("../sample.txt")
	checkErr1(err)

	// 리소스 해제
	defer file.Close()

	// 읽기 예제 1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	checkErr2(err)

	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) : ", fileInfo, "\n")
	fmt.Println("파일 이름(2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 완료(1) : %d bytes\n\n", ct1)
	//fmt.Println(fd1)	// 타입 변환 안한 경우
	fmt.Println(string(fd1)) // 타입 변환

	fmt.Println("=============================")

	// 읽기 예제 2. - 탐색, Seek(Offset)
	o1, err := file.Seek(20, 0) // 20 만큼 커서가 이동한 후,
	// 0 - 처음 위치, 1 - 현재 위치, 2 - 마지막 위치
	checkErr1(err)

	fd2 := make([]byte, 100)
	ct2, err := file.Read(fd2)
	checkErr1(err)

	fmt.Printf("읽기 작업 완료(2) : %d bytes, %d ret\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")
	fmt.Println("=============================")

	// 읽기 예제 3.
	o2, err := file.Seek(0, 0)
	checkErr1(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 1) // offset 위치부터 읽어 온다.
	checkErr1(err)

	fmt.Printf("읽기 작업 완료(3) : %d bytes, %d ret\n\n", ct3, o2)
	// fmt.Println(string(fd3), "\n")
	// fmt.Println("=============================")

}
