// Switch문(2)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 제어문 (조건문) - switch2

	// 난수값
	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50 미만")
	default:
		fmt.Println("i -> ", i, " 기본값")

	}
}
