// 데이터 타입 : numberic 연산(3)
package main

import (
	"math"
)

func main() {
	// 예제 1. 오버플로우 에러: 범위 초과
	var n1 uint8 = math.MaxUint8 + 1
	var n2 uint16 = math.MaxUint16 + 1
	var n3 uint32 = math.MaxUint32 + 1
	var n4 uint64 = math.MaxUint64 + 1

	// 예제 2. 오버플로우 에러: 범위 미만
	var n5 uint8 = math.MinUint8 - 1
	var n6 uint16 = math.MinUint16 - 1
	var n7 uint32 = math.MinUint32 - 1
	var n8 uint64 = math.MinUint64 - 1
}
