// Write File(1)
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
	// Write File
	// Create : 새 파일 작성 및  파일 열기
	// Close : 리소드 닫기
	// Write, WriteString, WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요!
	// https://golang.org/pkg/os

	// 파일 리소스 생성
	file, err := os.Create("test_write.txt")

	checkErr1(err)
	checkErr2(err)

	// 리소스 해제
	defer file.Close()

	// 쓰기 예제 1. - Write
	s1 := []byte{48, 49, 50, 101, 115}
	n1, err2 := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	checkErr2(err2)

	fmt.Printf("Write 1 completed : %d\n", n1)

	file.Sync() // Commit for Write

	// 쓰기 예제 2. - WriteString
	s2 := "Hello! golang!\n File Write Test! -1 \n"
	n2, err2 := file.WriteString(s2)
	checkErr2(err2)

	fmt.Printf("Write 2 completed : %d\n", n2)

	file.Sync() // Commit for Write

	// 쓰기 예제 3. - WriteAt
	s3 := "Test WriteAt! -2 \n"
	n3, err3 := file.WriteAt([]byte(s3), 70) // len(offset) 조절하면서 테스트
	checkErr2(err3)

	fmt.Printf("Write 3 completed : %d\n", n3)

	file.Sync() // Commit for Write

	// 쓰기 예제 4. - WriteString
	s4 := "Hello golang! \n File Write Test! - 3\n"
	n4, err4 := file.WriteString(s4)
	checkErr1(err4)

	fmt.Printf("Write 4 completed : %d\n", n4)
}
