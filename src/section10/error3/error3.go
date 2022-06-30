// Error handling (3)

package main

import (
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	_ "log"
	_ "os"
)

func main() {
	// Error handling
	// error 패키지의 New 메서드를 활용한 에러 생성
	// fmt.Errorof 보다 더 많이 사용

	var err1 error = errors.New("Error occured - 1")
	err2 := errors.New("Error occured - 2")

	fmt.Println("err1 : ", err1)
	fmt.Println("err1 : ", err1.Error())

	fmt.Println("err2 : ", err2)
	fmt.Println("err2 : ", err2.Error())
}
