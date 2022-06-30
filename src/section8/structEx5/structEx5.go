// 구조체 심화(5)
// 메소드 오버라이딩

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

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee             // is a 관계
	specialBonus float64 "특별보너스"
}

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴

	// 예제 1.
	// 직원
	emp1 := Employee{"kim", 10000000, 1500000}
	emp2 := Employee{"lee", 20000000, 2500000}

	// 임원
	ex1 := Executives{Employee{"lee", 30000000, 5000000}, 100000000}

	fmt.Println("emp1 : ", int(emp1.Calculate()))
	fmt.Println("emp2 : ", int(emp2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex1 : ", int(ex1.Calculate()+ex1.specialBonus))	// 오버라이딩으로 인한 잘못된 값 반환
	fmt.Println("ex1 : ", int(ex1.Calculate()))						// 오버라이딩으로 올바른 값 반환
	fmt.Println("ex1 : ", int(ex1.Employee.Calculate()+ex1.specialBonus))
}
