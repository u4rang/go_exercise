// Function(6)
package main

import (
	"fmt"
)

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// Function
	// Function 을 변수에 할당

	// 예제 1. Slice 에 할당
	f1 := []func(int, int) int{multiply, sum}

	a := f1[0](1, 2)
	b := f1[1](1, 2)

	fmt.Println("a :", a)
	fmt.Println("b :", b)

	// 예제 2. 변수에 할당
	var f2 func(int, int) int = multiply
	f3 := sum

	fmt.Println("f2(2, 3) :", f2(2, 3))
	fmt.Println("f3(2, 3) :", f3(2, 3))

	// 예제 3. Map 에 할당
	m := map[string]func(int, int) int{
		"mulFunc": multiply,
		"sumFunc": sum,
	}

	fmt.Println("m[\"mulFunc\"](3, 4) :", m["mulFunc"](3, 4))
	fmt.Println("m[\"sumFunc\"](3, 4) :", m["sumFunc"](3, 4))

}
