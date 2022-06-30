// Function(5)
package main

import (
	"fmt"
)

func multiply(n ...int) int {
	tot := 1

	for _, v := range n {
		tot *= v
	}

	return tot
}

func sum(n ...int) int {
	tot := 0

	for _, v := range n {
		tot += v
	}

	return tot
}

func prtWord(msg ...string) {
	for _, v := range msg {
		fmt.Println("prtWord : ", v)
	}
}

func main() {
	// Function
	// 가변 인자 (매개 변수 개수가 동적으로 변할 때, 정해져 있지 않음)

	// 예제 1.
	x := multiply(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("multiply : ", x)

	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("sum : ", y)

	// 예제 2.
	prtWord("Hello", "World", "golang")

	// 예제 3.
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	m := multiply(a...)
	n := sum(a...)

	fmt.Println("multiply with array : ", m)
	fmt.Println("sum with array : ", n)
}
