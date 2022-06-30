// Error handling (1)

package main

import (
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	"os"
	_ "os"
)

func main() {
	// Error handling
	// - 소프트웨어의 품질 향상 가장 중요한 것  !
	// - 유형 코드 및 에러 정보 등 정보를 남기는 것!
	// - golang 에서는 기본적으로 error 패키지에서 error 인터페이스를 제공
	// type error interface { Error() string}
	// 즉, Error 메소드를 구현하면 사용자 정의 에러 처리 제작 가능
	// 기본적으로 메소드마다 리턴 타입 2개(리턴값, 에러)
	// 주로 1. Errorf(에러 타입 리턴) 메소드, 2. Fatal(프로그램 종료) 메서드 통해 에러 출력

	// 기본적인 메서드 에러 처리 예제
	f, error := os.Open("unnamedfile") // 일부러 에러 발생
	// f, error := os.Open("/Users/chaesystersfather/Dev/go_exercise/README.md")
	if error != nil {
		log.Fatal(error.Error()) // 방법 1.
		// log.Fatal(error) // 방법 2.
	}
	fmt.Println("--------")
	fmt.Println("f Name : ", f.Name())

}
