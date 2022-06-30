// 열거형3
package main

import "fmt"

func main() {
	// 열거형
	// _ 로 생략 가능
	const (
		_ = iota
		A
		B
		_
		C
	)

	fmt.Println(A, B, C)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILEVER
		_
		GOLD
		PLATINUM
	)

	fmt.Println(DEFAULT, SILEVER, GOLD, PLATINUM)
}
