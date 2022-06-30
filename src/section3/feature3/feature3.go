// GO 특징(3)

package main

import "fmt"

func main() {
	// GO의 fmt 기능 - 강제 포메팅, 코드 서식에 대한 강제
	// 한사람이 코딩 한 것 같은 일관성 유지
	// 코드 스타일 유지
	// gofmt -h : 사용법
	// gofmt -w : 원본파일에 반영

	// 예제 1.
	// for i:=0; i <= 10; i++ {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
