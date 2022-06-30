// 구조체 심화(4)
// - 상속

package main

import (
	"fmt"
	_ "fmt"
)

type Employee struct {
	name   string  "직원이름"
	salary float64 "급여"
	bonus  float64 "보너스"
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee             // is a 관계
	specialBonus float64 "특별보너스"
}

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재 사용하는 관점 제공
	// 상속을 허용하지 않는 golang 언어에서 메소드 상속 활용을 위한 패턴

	// 예제 1.
	// 직원
	emp1 := Employee{"kim", 10000000, 1500000}
	emp2 := Employee{"lee", 20000000, 2500000}

	// 임원
	ex1 := Executives{Employee{"lee", 30000000, 5000000}, 100000000}

	fmt.Println("emp1 : ", int(emp1.Calculate()))
	fmt.Println("emp2 : ", int(emp2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex1 : ", int(ex1.Calculate()+ex1.specialBonus))
}
