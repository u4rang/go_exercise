// 자료형 : slice(5)
package main

import (
	"fmt"
)

func main() {
	// Slice 복사
	// copy(복사 대상, 원본)
	// - 반드시, make로 공간 할당 후 복사 필요
	// - 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	// 예제 1. 복사
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Printf("slice1 len : %d, cap : %d, %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2 len : %d, cap : %d, %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("slice3 len : %d, cap : %d, %v\n", len(slice3), cap(slice3), slice3)

	// 예제 2. 원본에 영향 없음
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 6
	b[4] = 10

	fmt.Printf("a len : %d, cap : %d, %v\n", len(a), cap(a), a)
	fmt.Printf("b len : %d, cap : %d, %v\n", len(b), cap(b), b)

	// 예제 3. 추출은 참조
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조 -> 원본 값 변경 된다.

	d[1] = 7

	fmt.Printf("c len : %d, cap : %d, %v\n", len(c), cap(c), c)
	fmt.Printf("d len : %d, cap : %d, %v\n", len(d), cap(d), d)

	// 예제 4. 추출하면서 용량 지정
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]

	f[1] = 12

	fmt.Printf("e len : %d, cap : %d, %v\n", len(e), cap(e), e)
	fmt.Printf("f len : %d, cap : %d, %v\n", len(f), cap(f), f)

}
