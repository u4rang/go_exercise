// Function(7)
package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("prtHello : (", n, ")", "hi")
	prtHello(n - 1)
}

func main() {
	// Function
	// 재귀 함수 (Recursion)
	// 장점
	//  - 프로그램이 보기 쉽고
	//  - 코드 간결
	//  - 오류 수정 용이
	// 단점
	//  - 코드 이해 어려움
	//  - 기억 공간을 많이 사용
	//  - 무한 루프 가능성

	// 예제 1.
	x := fact(10)
	fmt.Println("x :", x)

	// 예제 2.
	prtHello(10)
}
