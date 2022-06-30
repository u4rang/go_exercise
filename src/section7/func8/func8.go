// Function(8)
package main

import (
	"fmt"
)

func main() {
	// Function
	// 익명 함수, Annonymous Function
	// 즉시 실행 (선언과 동시에)

	// 예제 1.
	func() {
		fmt.Println("Annonymous Function 1")
	}()

	// 예제 2.
	msg := "Hello golang!"

	func(m string) {
		fmt.Println("msg : ", m)
	}(msg)

	// 예제 3.
	func(x, y int) {
		fmt.Println("x + y = ", x+y)
	}(10, 20)

	// 예제 4.
	r := func(x, y int) int {
		return x + y
	}(10, 100)

	fmt.Println("r : ", r)
}
