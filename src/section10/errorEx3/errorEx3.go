// Error handling 심화(3)

package main

import (
	_ "errors"
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	"math"
	_ "math"
	"time"
)

// 예외(에러) 처리 구조체
type PowError struct {
	time    time.Time   // 에러 발생 시간
	value   interface{} // Parameter
	message string      // Error message
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %f) - %s\n", e.time, e.value.(float64), e.message)
}

// f의 i 제곱 함수
func Pow(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "Zero is not allowed"}
	}

	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "String is not allowed"}
	}

	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: i, message: "String is not allowed"}
	}

	return math.Pow(f, i), nil // 제곱, nil 반환
}

func main() {
	// Error handling 심화
	// error 타입이 아닌 경우 에러 처리 방법
	// Error 메소드를 구현해서 사용자 정의 에러 심화 처리 방법
	// 구조체를 사용해서 세부적인 정보 출력

	fmt.Println("Start!")

	// 정상
	if r1, e1 := Pow(2, 4); e1 == nil {
		fmt.Println("Pow(2, 4) = ", r1)
	} else {
		log.Fatal(e1)
		// log.Fatal(e1.Error())
	}

	// 예외 발생
	if r2, e2 := Pow(0, 4); e2 == nil {
		fmt.Println("Pow(0, 4) = ", r2)
	} else {
		log.Fatal(e2)
		// log.Fatal(e2.Error())
	}

	// 예외 발생
	// if r3, e3 := Pow("Hi", 4); e3 == nil {
	// 	fmt.Println("Pow(\"Hi\", 4) = ", r3)
	// } else {
	// 	log.Fatal(e3)
	// 	// log.Fatal(e3.Error())
	// }
}
