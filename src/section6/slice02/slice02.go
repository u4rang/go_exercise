// 자료형 : slice(2)
package main

import (
	"fmt"
)

func main() {
	// Slice 참조 타입

	// 예제 1. 배열
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 // 복사

	arr2[0] = 7
	fmt.Printf("arr1 : %v\n", arr1)
	fmt.Printf("arr2 : %v\n", arr2)

	// 예제 2. Slice
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1 // 참조

	slice2[0] = 7

	fmt.Printf("slice1 : %v\n", slice1)
	fmt.Printf("slice2 : %v\n", slice2)

	// 예제 3. Slice 예외 상황
	slice3 := make([]int, 50, 100) // 초기화는 데이터 값의 기본값
	fmt.Println("slice3[4] : ", slice3[4])
	// fmt.Println("slice3[60] : ", slice3[60]) // Index Over 예외

	// 예제 4. Slice 순회
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}
