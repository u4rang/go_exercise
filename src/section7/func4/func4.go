// Function(4)
package main

import (
	"fmt"
)

func multiply(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 10

	return r1, r2
}

func multiply2(x, y int) (int, int) {
	r1 := x * 10
	r2 := y * 10

	return r1, r2
}

func multiply3(x, y int) (int, int) {
	return x * 10, y * 10
}

func main() {
	// Function
	// 리턴 값 변수 사용
	a, b := multiply(10, 100)
	fmt.Println("a :", a, ", b :", b)

	c, d := multiply2(10, 100)
	fmt.Println("c :", c, ", d :", d)

	e, f := multiply3(10, 100)
	fmt.Println("e :", e, ", f :", f)
}
