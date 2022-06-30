// 자료형 : 배열(2)
package main

import (
	"fmt"
)

func main() {
	// 배열 순회

	// 예제 1.
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	// Len 길이 반복
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Println()

	// 예제 2.
	arr2 := [5]int{7, 77, 777, 7777, 77777}

	// range - 더 많이 쓰임
	for i, v := range arr2 {
		fmt.Printf("[%d] - %d, ", i, v)
	}
	fmt.Println()

	// range - 인덱스 생략01
	for _, v := range arr2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// range - 인덱스 생략02
	for i := range arr2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
