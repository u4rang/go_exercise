// 데이터 타입 : numberic 연산(1)
package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산(산술, 비교)
	// 타입이 같아야 가능
	// 다른 타입 끼리는 반드시 형 변환 후 연산
	// 형 변환 없을 경우 예외(에러) 발생
	// +, -, *, /, %, <<, >>, &, ^

	// 예제 1.
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("uint8 : ", n1)
	fmt.Println("uint16 : ", n2)
	fmt.Println("uint32 : ", n3)
	fmt.Println("uint64 : ", n4)
	fmt.Println("MaxInt8 : ", math.MaxInt8)
	fmt.Println("MaxInt16 : ", math.MaxInt16)
	fmt.Println("MaxInt32 : ", math.MaxInt32)
	fmt.Println("MaxInt64 : ", math.MaxInt64)
	fmt.Println("MaxFloat32 : ", math.MaxFloat32)
	fmt.Println("MaxFloat64 : ", math.MaxFloat64)

	n5 := 100000 // int
	n6 := int16(1000)
	n7 := uint8(100)

	// fmt.Println("n5 + n6", n5+n6)	// 예외 발생
	fmt.Println("n5 + n6 :", n5+int(n6))
	fmt.Println("n6 + int16(n7) : ", n6+int16(n7))
	fmt.Println("n6 > int16(n7) : ", n6 > int16(n7))
	fmt.Println("n5 - int(n7) > 5000 : ", n5-int(n7) > 5000)
	fmt.Println("n7 : ", n7)

}
