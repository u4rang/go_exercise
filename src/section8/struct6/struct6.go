// 구조체(6)

package main

import (
	"fmt"
)

type Car struct { // 첫 글자 대문자로 선언
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct { // 첫 글자 소문자로 선언
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 중첩 구조체
	// - Car(spec)

	car1 := Car{
		"520d",
		"whilte",
		"bmw",
		spec{1000, 2000, 3000},
	}

	// 내부 spec 구조체 값 출력
	fmt.Println("car1.name : " + car1.name)
	fmt.Println("car1.color : " + car1.color)
	fmt.Println("car1.company : " + car1.company)
	fmt.Printf("car1.detail : %#v\n", car1.detail)

	// 예제 2.
	// 내부 spec 구조체 필드 값 출력
	fmt.Println("car1.detail.length", car1.detail.length)
	fmt.Println("car1.detail.height", car1.detail.height)
	fmt.Println("car1.detail.width", car1.detail.width)
}
