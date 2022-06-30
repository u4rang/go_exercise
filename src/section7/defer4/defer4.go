// Defer(4)
package main

import (
	"fmt"
)

func a() {
	defer end(start("defer in a")) // 중첩함수 사용 시, 주의!
	fmt.Println("There is in a")
}

func start(m string) string {
	fmt.Println("start : ", m)
	return m
}

func end(m string) {
	fmt.Println("end : ", m)
}

func main() {
	// Defer
	// - Function 안에서 Function 호출
	// 1. start
	// 2. There is in a
	// 3. end
	a()
}
