// Switch문(2)
package main

import (
	"fmt"
)

func main() {
	// 예제 1
	a := 30 / 15

	switch a {
	case 2, 4, 6, 8, 0:
		fmt.Println("a -> ", a, " 는 짝수")
	case 1, 3, 5, 7, 9:
		if a >= 20 {
			fmt.Println("a는 20보다 크다!")
		}
		fmt.Println("a -> ", a, " 는 홀수")
	}

	// 예제 2
	// - fallthrough
	// - 통과시키는 예제
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("Go!")
		fallthrough
	case "python":
		fmt.Println("Python!")
		fallthrough
	case "ruby":
		fmt.Println("Ruby!")
		// fallthrough
	case "c":
		fmt.Println("C!")
	}
}
