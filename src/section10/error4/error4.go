// Error handling (4)

package main

import (
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	_ "os"
)

func notZero(n int) (string, error) { // method 리턴 값 error 타입 중요!
	if n != 0 {
		s := fmt.Sprint("Hello golang :", n)
		return s, nil
	}

	return "", errors.New("0 is not a valid number.") // return value 으로 errors.New() 사용
}

func main() {
	// Error handling - 사용자 정의 에러 타입 리턴
	// Errorof 를 이용한 에러 처리

	a, err := notZero(1)

	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
	}

	fmt.Println("a : ", a)

	b, err := notZero(0)

	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
	}

	// Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("b : ", b)

	fmt.Println("End Error handling!")
}
