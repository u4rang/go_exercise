// 구조체 심화(1)

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

// 생성자 패턴
func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest} // 구조체 인스턴스를 생성 한 뒤, 포인터를 리턴
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제 1.
	kim := Account{number: "123-456-789", balance: 10000000, interest: 0.015}

	// 예제 2.
	var lee *Account = new(Account)
	lee.number = "456-789-123" // getter, setter
	lee.balance = 20000000
	lee.interest = 0.015

	fmt.Println("kim : ", kim)
	fmt.Println("lee : ", lee)

	park := NewAccount("789-123-456", 30000000, 0.035)

	fmt.Println("park : ", park)
}
