// 자료형 : pointer(1)
package main

import (
	"fmt"
)

func main() {
	// Pointer
	// Go : Pointer 지원(C)
	// 변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택
	// Python, Java(JRE), -> 컴파일러, 인터프리터
	// 포인터를 지원하지 않는 언어 - Python, Java, C#
	// 주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	// *(에스터리스크)로 사용
	// nil 로 초기화 (nil != 0)
	// & - 메모리 번지수 출력

	// 예제 1.
	var a *int            // 방법 1
	var b *int = new(int) // 방법 2

	// a = 8 // 에러 발생, 주소임

	fmt.Println("a : ", a) // &
	fmt.Println("b : ", b)

	i := 7

	a = &i // & 메모리 주소 전달
	b = &i

	fmt.Println("i : ", i, &i)
	fmt.Println("a : ", *a, a, &a) // *a - 역참조 - 번지수를 찾아가서 실제 값을 확인
	fmt.Println("b : ", *b, b, &b) // *b - 역참조 - 번지수를 찾아가서 실제 값을 확인

	var c = &i
	d := &i

	fmt.Println("c : ", *c, c, &c)
	fmt.Println("d : ", *d, d, &d)
	fmt.Println("----------------------")

	// 참조형 테스트
	*d = 10

	fmt.Println("i : ", i, &i)
	fmt.Println("a : ", *a, a, &a)
	fmt.Println("b : ", *b, b, &b)
	fmt.Println("c : ", *c, c, &c)
	fmt.Println("d : ", *d, d, &d)

}
