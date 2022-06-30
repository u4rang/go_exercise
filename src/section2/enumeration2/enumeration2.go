// 열거형2
package main

import "fmt"

func main() {
	const (
		A = (iota + 1) * 10
		B
		C
	)

	fmt.Println(A, B, C)
}
