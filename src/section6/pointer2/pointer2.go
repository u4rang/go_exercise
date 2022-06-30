// 자료형 : pointer(2)
package main

import (
	"fmt"
)

func main() {
	// Pointer
	// 예제 1.
	i := 7
	p := &i // Pointer 형

	fmt.Println("i : ", i)
	fmt.Println("&i : ", &i)
	fmt.Println("*p : ", *p) // 역참조
	fmt.Println("p : ", p)
	fmt.Println("-------------------")
	*p++ // 주소값을 연산한 것이 아니라서 정상적으로 동작
	fmt.Println("i : ", i)
	fmt.Println("&i : ", &i)
	fmt.Println("*p : ", *p) // 역참조
	fmt.Println("p : ", p)
	fmt.Println("-------------------")
	*p = 9999 // 포인터 변수 역 참조 값 변경
	fmt.Println("i : ", i)
	fmt.Println("&i : ", &i)
	fmt.Println("*p : ", *p) // 역참조
	fmt.Println("p : ", p)
	fmt.Println("-------------------")
	i = 7777
	fmt.Println("i : ", i)
	fmt.Println("&i : ", &i)
	fmt.Println("*p : ", *p) // 역참조
	fmt.Println("p : ", p)
}
