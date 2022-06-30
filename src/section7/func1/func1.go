// Function(1)
package main

import (
	"fmt"
	"strconv" // 숫자를 문자로 변환 할 때 사용하는 패키지
)

// func 함수명() : 매개변수 없음, 반환 값 없음
func helloGolang() {
	fmt.Println("helloGolang")
}

// func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
func sayOne(msg string) {
	fmt.Println("sayOne", msg)
}

// func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
func sum(x int, y int) int {
	return x + y
}

func main() {
	// Function
	// 선언: func 키워드로 선언 - 4가지
	// func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
	// func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재
	// func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
	// func 함수명() : 매개변수 없음, 반환 값 없음
	// 타 언어와 달리 반환 값(return value) 다수 가능
	// 함수 선언 위치는 어느 곳이든 가능

	// 예제 1.
	helloGolang()

	// 예제 2.
	sayOne("Hello world!")

	// 예제 3. - 반환 값 1개
	result := sum(1, 2)
	fmt.Println("sum : ", result)
	fmt.Println("sum : ", sum(1, 2))
	fmt.Println("sum : ", strconv.Itoa(sum(1, 2)))
}
