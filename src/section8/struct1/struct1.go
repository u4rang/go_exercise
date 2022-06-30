// 구조체(1)

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
	// golang 필드들의 집합체 또는 컨테이너
	// 필드는 갖지만 메소드는 갖지 않는다.
	// 객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	// 상속, 객체, 클래스 개념 없음
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	// 예제 1.
	kim := Account{number: "123-456-789", balance: 10000000, interest: 0.015}
	lee := Account{number: "987-654-321", balance: 10000000}
	park := Account{number: "456-789-234", interest: 0.025}
	cho := Account{"789-123-456", 15000000, 0.03}

	fmt.Println("kim :", kim)
	fmt.Println("lee :", lee)
	fmt.Println("park :", park)
	fmt.Println("cho :", cho)
	fmt.Println()

	fmt.Println("kim's interest :", int(kim.Calculate()))
	fmt.Println("lee's interest :", int(lee.Calculate()))
	fmt.Println("park's interest :", int(park.Calculate()))
	fmt.Println("cho's interest :", int(cho.Calculate()))

}
