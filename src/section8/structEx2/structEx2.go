// 구조체 심화(2)

package main

import (
	"fmt"
	_ "fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// 예제 1.
	kim := Account{"123-456-789", 10000000, 0.015}
	lee := Account{"456-789-123", 20000000, 0.025}

	fmt.Println("kim :", kim)
	fmt.Println("lee :", lee)

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println()

	fmt.Println("kim :", kim)
	fmt.Println("lee :", lee)

}
