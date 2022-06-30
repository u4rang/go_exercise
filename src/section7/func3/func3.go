// Function(3)
package main

import (
	"fmt"
)

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e, f int) (int, int, int, int, int, int) {
	return a * 10, b * 100, c * 1000, d * 10000, e * 100000, f * 100000
}

func main() {
	// Function
	// 다중 값 반환
	// 예제 1
	result1, result2 := multiply(10, 100)
	result3, _ := multiply(10, 100)
	_, result4 := multiply(10, 100)
	fmt.Println("result1 : ", result1, "result2 :", result2)
	fmt.Println("result3 : ", result3, "result3 :", result4)

	// 예제 2
	a, b, c, d, e, f := arrayMultiply(1, 2, 3, 4, 5, 6)
	fmt.Println("arrayMultiply", a, b, c, d, e, f)
}
