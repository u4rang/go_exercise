// Defer(1)
package main

import (
	"fmt"
)

func exF1() {
	fmt.Println("f1 Method Start!")
	defer exF2() // 마지막에 호출
	fmt.Println("f1 Method end!")
}

func exF2() {
	fmt.Println("f2 is called")
}

func main() {
	// Defer
	// - 함수 실행(지연)
	// - Defer 를 호출한 함수가 종료되기 직전에 호출된다.
	// - 타 언어의 Finally 문과 비슷
	// - 용도 : 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제, 변수 초기화

	// 예제 1.
	exF1()
}
