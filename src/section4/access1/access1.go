// 패키지 접근제어 (1)

package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	// 패키지 접근제어
	// 변수, 상수, 함수, 메소드, 구조체 등 식별자
	// 첫글자 기준
	// - 대문자 : 패키지 외부에서 접근 가능
	// - 소문자 : 패키지 외부에서 접근 불가

	fmt.Println("100 보다 큰 수 ? : ", lib2.CheckNum1(101))
	fmt.Println("1000 보다 큰 수 ? : ", lib2.CheckNum2(999))
}
