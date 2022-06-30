// 구조체 심화(3)

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

// 리시버
func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 구조체 인스턴스 값 변경 시-> 포인터 전달, 보통의 경우 -> 값 전달
	// 예제 1.
	kim := Account{"123-456-789", 10000000, 0.015}
	lee := Account{"456-789-123", 20000000, 0.025}

	fmt.Println("kim :", kim)
	fmt.Println("lee :", lee)

	kim.CalculateD(100000)
	lee.CalculateP(100000)

	fmt.Println()

	fmt.Println("kim :", int(kim.balance))
	fmt.Println("lee :", int(lee.balance))

}
