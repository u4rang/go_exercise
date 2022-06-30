// 데이터 타입 : numberic 연산(2)
package main

import (
	"fmt"
)

func main() {
	// 예제 1.
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("n1 + n2 : ", n1+n2)
	fmt.Println("n1 - n2 : ", n1-n2)
	fmt.Println("n1 * n2 : ", n1*n2)
	fmt.Println("n1 / n2 : ", n1/n2)
	fmt.Println("n1 % n2 : ", n1%n2)
	fmt.Println("n1 << 2 : ", n1<<2)
	fmt.Println("n1 >> 2 : ", n1>>2)
	fmt.Println("^n1  : ", ^n1)

	// 예제 2.
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// fmt.Println("n3 + n4 : ", n3+n4)
	fmt.Println("flaot32(n3) + n4 : ", float32(n3)+n4) // 주의
	fmt.Println("n3 + int(n4) : ", n3+int(n4))         // 주의
	fmt.Println("n5 + uint16(n6) :", n5+uint16(n6))    // 주의

}
