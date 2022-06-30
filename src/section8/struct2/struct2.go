// 구조체(2)

package main

import (
	"fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

// 리시버 형태로 클래스와 메소드 연결 지원
func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 구조체
	// 다양한 선언 방법
	// &struct, struct
	// - &struct : 포인터를 받아오기 역참조를 다시 하기 때문에 속도는 조금 느리다.
	//              인터페이스 메소를 선언 해둔 후 -> 오버라이딩 해서 메서드에 포인터 메서드를 사용할 경우 반드시 &struct 사용

	// 선언 방법 1. 포인터형
	var kim *Account = new(Account)
	kim.number = "123-456-789"
	kim.balance = 10000000
	kim.interest = 0.015

	// 선언 방법 2
	hong := &Account{number: "456-789-123", balance: 10000000, interest: 0.015}

	// 선언 방법 3
	lee := new(Account)
	lee.number = "789-123-456"
	lee.balance = 10000000
	lee.interest = 0.015

	fmt.Println("kim : ", kim)
	fmt.Println("hong : ", hong)
	fmt.Println("lee : ", lee)
	fmt.Println()

	fmt.Printf("kim : %#v\n", kim)
	fmt.Printf("hong : %#v\n", hong)
	fmt.Printf("lee : %#v\n", lee)
	fmt.Println()

	fmt.Println("kim.Calculate : ", kim.Calculate())
	fmt.Println("hong.Calculate : ", hong.Calculate())
	fmt.Println("lee.Calculate : ", lee.Calculate())

}
