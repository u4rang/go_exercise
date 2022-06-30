// 자료형 : pointer(3)
package main

import (
	"fmt"
)

func rptc(n *int) {
	*n = 77 // 역참조하여 값을 변경
}

func vptc(n int) {
	n = 66
}

func main() {
	// Pointer
	// Pointer에 값 전달
	// 함수, 메소드 호출 시에 매개변수 값을 복사하여 전달 -> 함수, 메소드 내에서는 원본 값 변경 불가능
	// 주로 원본 값 변경 위해서 Pointer로 전달
	// 특히 크기가 큰 배열인 경우, 값 복사 시에 시스템 부담 -> Pointer 전달로 해결(Slice, Map 참조 전달)

	// 예제 1.
	var a int = 10
	var b int = 10

	fmt.Println("a =", a, "b =", b)

	rptc(&a)
	vptc(b)

	fmt.Println("a =", a, "b =", b)
}
