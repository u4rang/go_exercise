// 인터페이스 고급(3)
package main

import (
	"fmt"
	_ "fmt"
)

/* 구조체에 대한 값 및 포인터 체크
func checkType(arg interface{}) {
	switch arg.(type) {
	case Account:
	case *Account:
	default:
	}
}
*/

func checkType(arg interface{}) {
	// arg.(type)을 통해서 현재 데이터 타입 반환
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int", arg)
	case float32, float64:
		fmt.Println("This is a float", arg)
	case string:
		fmt.Println("This is a string", arg)
	case nil:
		fmt.Println("This is a nil", arg)
	default:
		fmt.Println("What is this type?", arg)
	}
}

func main() {
	// 인터페이스 고급 - 타입 검사
	// 실제 타입 검사 시 switch 를 사용
	// 빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로,
	// 타입 체크를 통해 형 변환 후 사용 가능 - 런타임 에러 방지

	// 예제 1.
	checkType(true)
	checkType(1)
	checkType(1.23456)
	checkType("Hello, world!")
	checkType(nil)

}
