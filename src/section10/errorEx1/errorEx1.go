// Error handling 심화(1)

package main

import (
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	"math"
	_ "math"
)

// f의 i제곱 구하는 함수
func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0 is not a valid power number")
	}

	return math.Pow(f, i), nil // 제곱, nil 반환
}

func main() {
	// Error handling 심화
	// golang error 패키지의 New 메서드 사용 -> 사용자 에러 처리 생성

	// 예제 1.
	if f, err := Power(7, 3); err != nil {
		// fmt.Printf("Error Message : %s\n", err.Error())
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("Power( 7, 3) : ", f)
	}

	if f, err := Power(0, 3); err != nil {
		// fmt.Printf("Error Message : %s\n", err.Error())
		fmt.Printf("Error Message : %s\n", err)
	} else {
		fmt.Println("Power( 0, 3) : ", f)
	}
}
